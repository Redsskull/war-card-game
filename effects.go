package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// Clickable card widget with hover effects
type ClickableCard struct {
	widget.BaseWidget
	cardImage   *canvas.Image
	cardCount   *widget.Label
	hoverBorder *canvas.Rectangle
	onTap       func()
}

// NewClickableCard creates a new clickable card with hover effects
func NewClickableCard(cardImage *canvas.Image, cardCount *widget.Label, onTap func()) *ClickableCard {
	card := &ClickableCard{
		cardImage: cardImage,
		cardCount: cardCount,
		onTap:     onTap,
	}

	// Create hover border (hidden initially)
	card.hoverBorder = canvas.NewRectangle(color.RGBA{255, 215, 0, 150}) // Golden glow
	card.hoverBorder.StrokeColor = color.RGBA{255, 215, 0, 255}
	card.hoverBorder.StrokeWidth = 2
	card.hoverBorder.Hide()

	card.ExtendBaseWidget(card)
	return card
}

// Implement Tappable interface
func (c *ClickableCard) Tapped(*fyne.PointEvent) {
	if c.onTap != nil {
		c.onTap()
	}
}

// Implement Hoverable interface
func (c *ClickableCard) MouseIn(*desktop.MouseEvent) {
	c.hoverBorder.Show()
	c.hoverBorder.Refresh()
}

func (c *ClickableCard) MouseOut() {
	c.hoverBorder.Hide()
	c.hoverBorder.Refresh()
}

// CreateRenderer creates the visual representation
func (c *ClickableCard) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(
		container.NewStack(c.hoverBorder, c.cardImage, c.cardCount))
}
