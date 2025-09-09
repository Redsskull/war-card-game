package main

import (
	"fmt"
)

// Card represents a playing card with a number and suit.
type Card struct {
	Value int    // 2-18 (Joker = 15 Ace = 14, King = 13, Queen = 12, Jack = 11)
	Suit  string // Hearts, Diamonds, Clubs, Spades and so on
}

// String method
func (c Card) String() string {
	return fmt.Sprintf("%d of %s", c.Value, c.Suit)
}
