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

	// Test Player
	fmt.Println("\n=== Testing Player ===")

	// Create a player
	player1 := Player{Name: "Test Player", Cards: []Card{}}

	// Give them a few cards
	testCard1 := Card{Value: 5, Suit: "Hearts"}
	testCard2 := Card{Value: 10, Suit: "Spades"}
	testCard3 := Card{Value: 14, Suit: "Diamonds"}

	player1.AddCard(testCard1)
	player1.AddCard(testCard2)
	player1.AddCard(testCard3)

	fmt.Println("Player after adding cards:", player1)
	fmt.Println("Player has cards?", player1.HasCards())

	// Play some cards
	fmt.Println("\nPlaying cards:")
	card1 := player1.PlayCard()
	fmt.Println("Played:", card1)
	fmt.Println("Player now:", player1)

	card2 := player1.PlayCard()
	fmt.Println("Played:", card2)
	fmt.Println("Player now:", player1)

	card3 := player1.PlayCard()
	fmt.Println("Played:", card3)
	fmt.Println("Player now:", player1)
	fmt.Println("Player has cards?", player1.HasCards())

	// Try to play when empty
	emptyCard := player1.PlayCard()
	fmt.Println("Tried to play when empty, got:", emptyCard)

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
