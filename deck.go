package main

import "math/rand"

type Deck struct {
	Cards []Card // Slice of Card structs
}

// NewDeck creates a full deck of 56 cards
func NewDeck() Deck {
	var cards []Card

	// Create regular cards 2-14 for each suit
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	for _, suit := range suits {
		for value := 2; value <= 14; value++ {
			card := Card{Value: value, Suit: suit}
			cards = append(cards, card)
		}
	}

	// Add special cards
	superColorfulAce := Card{Value: 15, Suit: "ColorfulBigAce"}
	superBlackAce := Card{Value: 16, Suit: "BlackBigAce"}
	colorfulJoker := Card{Value: 17, Suit: "ColorfulJoker"}
	blackJoker := Card{Value: 18, Suit: "BlackJoker"}

	cards = append(cards, superColorfulAce, superBlackAce, colorfulJoker, blackJoker)

	return Deck{Cards: cards}
}

// Shuffle randomizes the order of cards in the deck
// Fisher-Yates shuffle algorithm
func (d *Deck) Shuffle() {
	for i := len(d.Cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}
