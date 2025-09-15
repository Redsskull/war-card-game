package main

import (
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
