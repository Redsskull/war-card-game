package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Initialize the game
	player1, cpu := StartGame()

	myApp := app.New()
	myWindow := myApp.NewWindow("War Card Game")
	myWindow.Resize(fyne.NewSize(800, 600))

	// Title
	title := widget.NewLabel("🃏 War Card Game")
	title.TextStyle.Bold = true

	// Card displays - make them BIG and readable
	playerCardDisplay := widget.NewLabel("Ready to play!")
	cpuCardDisplay := widget.NewLabel("Ready to play!")

	// Make text bigger and centered
	playerCardDisplay.Alignment = fyne.TextAlignCenter
	cpuCardDisplay.Alignment = fyne.TextAlignCenter

	// Score displays - show ACTUAL card counts
	playerScore := widget.NewLabel(fmt.Sprintf("Your cards: %d", len(player1.Cards)))
	cpuScore := widget.NewLabel(fmt.Sprintf("CPU cards: %d", len(cpu.Cards)))

	updateScores := func() {
		playerScore.SetText(fmt.Sprintf("Your cards: %d", len(player1.Cards)))
		cpuScore.SetText(fmt.Sprintf("CPU cards: %d", len(cpu.Cards)))
	}

	// Game result - simple and clean
	gameResult := widget.NewLabel("Click 'Play Round' to start!")
	gameResult.Alignment = fyne.TextAlignCenter

	// Simple play button
	var playButton *widget.Button
	playButton = widget.NewButton("Start Game", func() {
		if !player1.HasCards() || !cpu.HasCards() {
			return
		}
		if playButton.Text == "Start Game" {
			playButton.SetText("Play Round")
		}

		// Get the cards and result
		playerCard, cpuCard, result := PlayRound(player1, cpu)

		// 🎯 Show the cards clearly and big!
		playerCardDisplay.SetText(fmt.Sprintf("YOU PLAYED:\n%s", playerCard.String()))
		cpuCardDisplay.SetText(fmt.Sprintf("CPU PLAYED:\n%s", cpuCard.String()))

		// Show result cleanly
		gameResult.SetText(result)
		updateScores()

		// Check game over
		if !player1.HasCards() {
			gameResult.SetText("GAME OVER!\nCPU WINS! 🤖")
			playButton.SetText("Game Over")
			playButton.Disable()
		} else if !cpu.HasCards() {
			gameResult.SetText("GAME OVER!\nYOU WIN! 🏆")
			playButton.SetText("Game Over")
			playButton.Disable()
		}
	})

	// Layout section

	// TOP: Title section
	titleSection := container.NewCenter(
		widget.NewLabel("🃏 War Card Game"))

	// CPU section (under title)
	cpuSection := container.NewCenter(cpuScore)
	// TODO: This will replace the title once game starts

	// TOP AREA: Title + CPU together
	topArea := container.NewVBox(
		titleSection,
		widget.NewSeparator(),
		cpuSection)

	// MIDDLE: Game results (later card icons)
	middleArea := container.NewCenter(gameResult)

	// BOTTOM: Score left, button right
	bottomArea := container.NewCenter(
		container.NewHBox(
			playerScore, // Left side: "Your cards: 28"
			playButton)) // Right side: "Start Game" button

	// 🎯 USE BORDER LAYOUT: Distributes across full screen
	content := container.NewBorder(
		topArea,    // Top of screen
		bottomArea, // Bottom of screen
		nil,        // Left side (none)
		nil,        // Right side (none)
		middleArea) // Center/middle of screen

	// Add light blue background
	background := canvas.NewRectangle(color.RGBA{70, 130, 180, 255}) // blue
	contentWithBackground := container.NewStack(background, container.NewPadded(content))
	myWindow.SetContent(contentWithBackground)

	myWindow.ShowAndRun()
}
