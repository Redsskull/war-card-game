package main

import (
	"strings"
	"testing"
)

func TestStartGame(t *testing.T) {
	player1, cpu, messages := StartGame()

	// Test that players are created
	if player1 == nil || cpu == nil {
		t.Error("Players should not be nil")
	}

	// Test that cards were dealt
	totalCards := len(player1.Cards) + len(cpu.Cards)
	if totalCards != 55 { // Your deck has 55 cards
		t.Errorf("Expected 55 total cards, got %d", totalCards)
	}

	// Test that messages were created
	if len(messages) == 0 {
		t.Error("Should have setup messages")
	}
}

func TestPlayRound(t *testing.T) {
	// Create test players with specific cards
	player1 := &Player{Name: "Test Player", Cards: []Card{
		{Value: 10, Suit: "Hearts"},
	}}
	cpu := &Player{Name: "Test CPU", Cards: []Card{
		{Value: 5, Suit: "Spades"},
	}}

	card1, card2, result := PlayRound(player1, cpu)

	// Player should win (10 > 5)
	if card1.Value != 10 || card2.Value != 5 {
		t.Error("Cards not played correctly")
	}

	// Test the result string contains expected info
	if !strings.Contains(result, "Player wins!") {
		t.Errorf("Expected result to contain 'Player wins!', got: %s", result)
	}

	if len(player1.Cards) != 2 { // Should have won both cards
		t.Errorf("Player should have 2 cards after winning, got %d", len(player1.Cards))
	}
}
