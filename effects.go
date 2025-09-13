package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// TODO: make this a canvs.text so it can have resizable text(maybe even moveable)
func cardBackWithCount(img *canvas.Image, count *widget.Label) fyne.CanvasObject {
	// Create a styled count with background for better visibility
	count.TextStyle.Bold = true

	// Create a container that positions the count at the bottom
	overlay := container.NewCenter(
		container.NewVBox(count),
		layout.NewSpacer(), // This pushes everything below to the bottom

	)

	return container.NewStack(img, overlay)
}

// ClickableCard widget: card image with anchored counter and click action
type ClickableCard struct {
	widget.BaseWidget
	cardImage *canvas.Image
	cardCount *widget.Label
	onTap     func()
}

func NewClickableCard(cardImage *canvas.Image, cardCount *widget.Label, onTap func()) *ClickableCard {
	card := &ClickableCard{
		cardImage: cardImage,
		cardCount: cardCount,
		onTap:     onTap,
	}
	card.ExtendBaseWidget(card)
	return card
}

func (c *ClickableCard) Tapped(*fyne.PointEvent) {
	if c.onTap != nil {
		c.onTap()
	}
}

// Renderer: image + counter
func (c *ClickableCard) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(
		cardBackWithCount(c.cardImage, c.cardCount),
	)
}
