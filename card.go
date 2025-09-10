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

// GetImageFilename returns the correct filename for this card's image
func (c Card) GetImageFilename() string {
	// Handle the 3 special jokers first
	switch c.Suit {
	case "Joker":
		return "Cards/card_joker.png" // Normal joker (15)
	case "Red Joker":
		return "Cards/card_joker_red.png" // Red joker (16)
	case "Black Joker":
		return "Cards/card_joker_black.png" // Black joker (17)
	}

	// Handle regular cards (2-14)
	// Convert suit name to match your filenames
	var suitName string
	switch c.Suit {
	case "Hearts":
		suitName = "heart" // Your files: "card_heart_..."
	case "Diamonds":
		suitName = "diamond" // Your files: "card_diamond_..."
	case "Clubs":
		suitName = "clubs" // Your files: "card_clubs_..."
	case "Spades":
		suitName = "spade" // Your files: "card_spade_..."
	default:
		return "Cards/card_joker.png" // Fallback
	}

	// Convert value to match your filenames
	var valueStr string
	if c.Value == 14 {
		valueStr = "ace" // Your files use "card_heart_ace.png"
	} else {
		valueStr = fmt.Sprintf("%d", c.Value) // "2", "3", "4"... "13"
	}

	return fmt.Sprintf("Cards/card_%s_%s.png", suitName, valueStr)
}
