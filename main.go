package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//testing the card creation
	testCard := Card{Value: 14, Suit: "Hearts"}
	fmt.Println("Test card:", testCard)

	//testing the deck creating
	testDeck := NewDeck()
	fmt.Printf("Deck has %d cards\n", len(testDeck.Cards))
	fmt.Println("First card:", testDeck.Cards[0])
	fmt.Println("Last card:", testDeck.Cards[len(testDeck.Cards)-1])

	myApp := app.New()
	myWindow := myApp.NewWindow("War Card Game")
	myWindow.Resize(fyne.NewSize(800, 600))

	// Create a simple welcome message
	welcome := widget.NewLabel("Welcome to War Card Game!")
	welcome.Alignment = fyne.TextAlignCenter

	// Create a button for testing
	testButton := widget.NewButton("Deal Cards", func() {
		welcome.SetText("Cards dealt! (not really yet)")
	})

	// Put everything in a container
	content := container.NewVBox(
		welcome,
		testButton,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun() // This uses myApp indirectly
}
