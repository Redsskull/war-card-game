package main

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// SoundSystem manages all audio for the game
type SoundSystem struct {
	shuffleBuffer  *beep.Buffer
	playCardBuffer *beep.Buffer
	initialized    bool
}

// Global sound system instance
var gameSound *SoundSystem

// InitializeSoundSystem sets up the audio system for the entire game
func InitializeSoundSystem() error {
	gameSound = &SoundSystem{}

	// Load shuffle sound
	shuffleFile, err := os.Open("sounds/card_shuffle.mp3")
	if err != nil {
		log.Printf("Warning: Could not load shuffle sound: %v", err)
		return err
	}
	defer shuffleFile.Close()

	shuffleStreamer, format, err := mp3.Decode(shuffleFile)
	if err != nil {
		log.Printf("Warning: Could not decode shuffle sound: %v", err)
		return err
	}
	defer shuffleStreamer.Close()

	// Initialize speaker with the audio format (do this only once!)
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Printf("Warning: Could not initialize speaker: %v", err)
		return err
	}

	// Load shuffle sound into memory buffer for instant playback
	gameSound.shuffleBuffer = beep.NewBuffer(format)
	gameSound.shuffleBuffer.Append(shuffleStreamer)

	// Load play card sound
	playCardFile, err := os.Open("sounds/playcard.mp3")
	if err != nil {
		log.Printf("Warning: Could not load play card sound: %v", err)
		return err
	}
	defer playCardFile.Close()

	playCardStreamer, cardFormat, err := mp3.Decode(playCardFile)
	if err != nil {
		log.Printf("Warning: Could not decode play card sound: %v", err)
		return err
	}
	defer playCardStreamer.Close()

	// Handle different sample rates by resampling before buffering
	gameSound.playCardBuffer = beep.NewBuffer(format)
	if cardFormat.SampleRate != format.SampleRate {
		// Resample the audio to match the speaker's sample rate
		resampled := beep.Resample(4, cardFormat.SampleRate, format.SampleRate, playCardStreamer)
		gameSound.playCardBuffer.Append(resampled)
	} else {
		// Same sample rate, no resampling needed
		gameSound.playCardBuffer.Append(playCardStreamer)
	}

	gameSound.initialized = true
	log.Println("🔊 Sound system initialized successfully!")
	return nil
}

// PlayShuffleSound plays the card shuffle sound (non-blocking)
func PlayShuffleSound() {
	if gameSound == nil || !gameSound.initialized {
		return // Silently ignore if sound system not initialized
	}

	// Create a new streamer from the buffer (so we can play multiple times)
	shuffleStreamer := gameSound.shuffleBuffer.Streamer(0, gameSound.shuffleBuffer.Len())
	speaker.Play(shuffleStreamer)
}

// PlayCardClickSound plays the card click sound (non-blocking)
func PlayCardClickSound() {
	if gameSound == nil || !gameSound.initialized {
		return // Silently ignore if sound system not initialized
	}

	// Create a new streamer from the buffer (so we can play multiple times)
	cardStreamer := gameSound.playCardBuffer.Streamer(0, gameSound.playCardBuffer.Len())
	speaker.Play(cardStreamer)
}
