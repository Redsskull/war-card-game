package main

import (
	"math"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
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

// breakLongText breaks long text into multiple lines at natural break points
// I decided to do this since I may lack the knowledge or Fyne can't keep my UI elements to the right if the
// text get's too long for the result, and I can't use canvas for a new line
func breakLongText(text string, maxLineLength int) string {
	if len(text) <= maxLineLength {
		return text
	}

	// Split by " -> " first to find natural break points
	parts := strings.Split(text, " -> ")
	if len(parts) == 1 {
		// No natural break points, just return original
		return text
	}

	result := parts[0]
	currentLineLength := len(parts[0])

	for i := 1; i < len(parts); i++ {
		nextPart := " -> " + parts[i]

		// If adding this part would exceed the line length, start a new line
		if currentLineLength+len(nextPart) > maxLineLength {
			result += "\n" + nextPart[4:] // Remove " -> " from start of new line
			currentLineLength = len(nextPart) - 4
		} else {
			result += nextPart
			currentLineLength += len(nextPart)
		}
	}

	return result
}

// newLargeMultilineText creates a RichText widget with larger text size that supports newlines
func newLargeMultilineText(text string) *widget.RichText {
	segment := &widget.TextSegment{
		Text: text,
		Style: widget.RichTextStyle{
			Alignment: fyne.TextAlignCenter,
			SizeName:  theme.SizeNameHeadingText, // This gives larger text like TextSize = 20. had to look up and learn themes for Fyne
			TextStyle: fyne.TextStyle{Bold: true},
		},
	}

	richText := widget.NewRichText(segment)
	return richText
}
