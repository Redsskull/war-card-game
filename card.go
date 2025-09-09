package main

import (
	"fmt"
)

// Card represents a playing card with a number and suit.
type Card struct {
	Value int    // 2-18
	Suit  string // Hearts, Diamonds, Clubs, Spades, and your special suits
}

// GetDisplayValue returns the card value as it appears on cards
func (c Card) GetDisplayValue() string {
	switch c.Value {
	case 14:
		return "Ace" // Full name instead of "A"
	case 13:
		return "King" // Full name instead of "K"
	case 12:
		return "Queen" // Full name instead of "Q"
	case 11:
		return "Jack" // Full name instead of "J"
	case 15:
		return "Super Ace" // Super Colorful Ace
	case 16:
		return "Super Black Ace" // Super Black Ace
	case 17:
		return "Colorful Joker" // Colorful Joker
	case 18:
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
	case "ColorfulBigAce":
		return "🌈" // Rainbow for colorful
	case "BlackBigAce":
		return "⚫" // Black circle
	case "ColorfulJoker":
		return "🃏" // Joker emoji
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
	case "Clubs", "Spades", "BlackBigAce", "BlackJoker":
		return "black"
	case "ColorfulBigAce", "ColorfulJoker":
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
	case 18:
		return "🖤 BLACK JOKER (Highest!)"
	case 17:
		return "🃏 COLORFUL JOKER"
	case 16:
		return "⚫ SUPER BLACK ACE"
	case 15:
		return "🌈 SUPER COLORFUL ACE"
	default:
		return fmt.Sprintf("%s of %s", c.GetDisplayValue(), c.Suit)
	}
}
