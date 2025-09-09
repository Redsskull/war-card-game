package main

import (
	"fmt"
)

// Player represents a player in the game.
type Player struct {
	Name  string // Name of the player
	Cards []Card // Cards held by the player
}

// PlayCard removes and returns the top card from player's hand
func (p *Player) PlayCard() Card {
	if len(p.Cards) == 0 {
		// Player has no cards left - they've lost!
		return Card{} // Empty card means "no card to play"
	}

	topCard := p.Cards[0] // Get first card
	p.Cards = p.Cards[1:] /// Create new slice starting from index 1
	return topCard
}

// AddCard adds a card to the bottom of player's hand
func (p *Player) AddCard(card Card) {
	p.Cards = append(p.Cards, card)
}

// HasCards checks if player still has cards
func (p *Player) HasCards() bool {
	return len(p.Cards) > 0
}

// String method to display player info nicely
func (p Player) String() string {
	return fmt.Sprintf("%s has %d cards", p.Name, len(p.Cards))
}
