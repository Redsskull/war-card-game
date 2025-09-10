package main

import (
	"fmt"
)

// Card represents a playing card with a number and suit.
type Card struct {
	Value int    // 2-17
	Suit  string // Hearts, Diamonds, Clubs, Spades, and your special suits
}

// GetDisplayValue returns the card value as it appears on cards
func (c Card) GetDisplayValue() string {
	switch c.Value {
	case 14:
		return "Ace"
	case 13:
		return "King"
	case 12:
		return "Queen"
	case 11:
		return "Jack"
	case 15:
		return "Joker" // Normal Joker
	case 16:
		return "Red Joker" // Red Joker
	case 17:
		return "Black Joker" // Black Joker
	default:
		return fmt.Sprintf("%d", c.Value)
	}
}

// GetSuitSymbol returns the symbol for the suit
func (c Card) GetSuitSymbol() string {
	switch c.Suit {
	case "Hearts":
		return "♥"
	case "Diamonds":
		return "♦"
	case "Clubs":
		return "♣"
	case "Spades":
		return "♠"
	case "Joker":
		return "🌈" // Rainbow for colorful
	case "RedJoker":
		return "🔴" // Red circle
	case "BlackJoker":
		return "🖤" // Black heart
	default:
		return "?"
	}
}

// GetSuitColor returns the color for styling
func (c Card) GetSuitColor() string {
	switch c.Suit {
	case "Hearts", "Diamonds":
		return "red"
	case "Clubs", "Spades", "BlackJoker":
		return "black"
	case "Red Joker", "Joker":
		return "rainbow" // I'll handle this specially in the UI
	default:
		return "black"
	}
}

// 🃏 Visual card display
func (c Card) GetCardDisplay() string {
	symbol := c.GetSuitSymbol()
	value := c.GetDisplayValue()

	// Special formatting for premium cards
	if c.Value >= 15 {
		return fmt.Sprintf("╔═════╗\n║%s%-4s║\n║  %s  ║\n║%-4s%s║\n╚═════╝",
			value, "", symbol, "", value)
	}

	// Regular cards
	return fmt.Sprintf("┌─────┐\n│%s%-4s│\n│  %s  │\n│%-4s%s│\n└─────┘",
		value, "", symbol, "", value)
}

// Compact display for in-game use
func (c Card) GetCompactDisplay() string {
	return fmt.Sprintf("[%s%s]", c.GetDisplayValue(), c.GetSuitSymbol())
}

// Enhanced description that shows the power hierarchy
func (c Card) String() string {
	switch c.Value {

	case 17:
		return "🃏 Black JOKER"
	case 16:
		return "⚫ Red Joker!"
	case 15:
		return "🌈 Joker!"
	default:
		return fmt.Sprintf("%s of %s", c.GetDisplayValue(), c.Suit)
	}
}
