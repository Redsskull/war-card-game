package main

import (
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// ClickableCard widget
type ClickableCard struct {
	widget.BaseWidget
	cardImage *canvas.Image
	onTap     func()
}

// NewClickableCard creates a new ClickableCard widget
func NewClickableCard(cardImage *canvas.Image, onTap func()) *ClickableCard {
	card := &ClickableCard{
		cardImage: cardImage,
		onTap:     onTap,
	}
	card.ExtendBaseWidget(card)
	return card
}

// Tapped handles the tap event
func (c *ClickableCard) Tapped(*fyne.PointEvent) {
	if c.onTap != nil {
		c.onTap()
	}
}

// Renderer: just the image
func (c *ClickableCard) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(c.cardImage)
}

// SHAKE ANIMATION

// ShakeContainer creates a dramatic shake effect for wars
func ShakeContainer(container *fyne.Container, duration time.Duration) {
	// Safety check - make sure container exists
	if container == nil {
		return
	}

	// Remember where the container started
	originalPos := container.Position()
	shakeIntensity := float32(45) // How many pixels to shake (adjust this!)

	// Create the animation
	shake := fyne.NewAnimation(duration, func(progress float32) {
		// progress goes from 0.0 to 1.0 as animation plays

		// Create realistic shake using sine waves
		// Fast shake: quick back-and-forth motion
		fastShake := float32(math.Sin(float64(progress * 30)))
		// Slow shake: slower, bigger movements
		slowShake := float32(math.Sin(float64(progress * 8)))

		// Make shake fade out over time (decay)
		decay := 1.0 - progress // Starts at 1.0, ends at 0.0

		// Combine shakes and apply decay
		offsetX := (fastShake*0.7 + slowShake*0.3) * shakeIntensity * decay
		offsetY := (fastShake*0.3 + slowShake*0.7) * shakeIntensity * decay * 0.5 // Less vertical

		//Move the container to new position
		newPos := fyne.NewPos(originalPos.X+offsetX, originalPos.Y+offsetY)
		container.Move(newPos)
	})

	// Make the animation smooth
	shake.Curve = fyne.AnimationEaseOut
	shake.Start()

	// Guarantee end at original position
	time.AfterFunc(duration, func() {
		fyne.Do(func() {
			container.Move(originalPos)
		})
	})
}
