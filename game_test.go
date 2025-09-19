package main

import (
	"strings"
	"testing"
)

func TestStartGame(t *testing.T) {
	player1, cpu := StartGame()

	// Test that players are created
	if player1 == nil || cpu == nil {
		t.Error("Players should not be nil")
	}

	// Test that cards were dealt
	totalCards := len(player1.Cards) + len(cpu.Cards)
	if totalCards != 55 { // Your deck has 55 cards
		t.Errorf("Expected 55 total cards, got %d", totalCards)
	}

	// Test that cards are distributed (each player should have cards)
	if len(player1.Cards) == 0 || len(cpu.Cards) == 0 {
		t.Error("Both players should have cards")
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

	card1, card2, result, warInfo := PlayRound(player1, cpu)

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

	// Test that it's not a war
	if warInfo.IsWar {
		t.Error("This should not be a war")
	}

	if warInfo.CardsAtStake != 2 {
		t.Errorf("Expected 2 cards at stake, got %d", warInfo.CardsAtStake)
	}
}

func TestWarRound(t *testing.T) {
	// Create test players with tied cards followed by different cards
	player1 := &Player{Name: "Test Player", Cards: []Card{
		{Value: 7, Suit: "Hearts"},  // Tied card
		{Value: 2, Suit: "Hearts"},  // War card 1
		{Value: 3, Suit: "Hearts"},  // War card 2
		{Value: 4, Suit: "Hearts"},  // War card 3
		{Value: 10, Suit: "Hearts"}, // Final deciding card
	}}
	cpu := &Player{Name: "Test CPU", Cards: []Card{
		{Value: 7, Suit: "Spades"}, // Tied card (same value!)
		{Value: 6, Suit: "Spades"}, // War card 1
		{Value: 8, Suit: "Spades"}, // War card 2
		{Value: 9, Suit: "Spades"}, // War card 3
		{Value: 5, Suit: "Spades"}, // Final deciding card
	}}

	card1, card2, result, warInfo := PlayRound(player1, cpu)

	// Test that it's a war
	if !warInfo.IsWar {
		t.Error("This should be a war")
	}

	// Test tied cards are stored correctly
	if warInfo.TiedCard1.Value != 7 || warInfo.TiedCard2.Value != 7 {
		t.Errorf("Expected tied cards to be 7 and 7, got %d and %d",
			warInfo.TiedCard1.Value, warInfo.TiedCard2.Value)
	}

	// Test final deciding cards
	if card1.Value != 10 || card2.Value != 5 {
		t.Errorf("Expected final cards to be 10 and 5, got %d and %d", card1.Value, card2.Value)
	}

	// Test that result contains "WAR!"
	if !strings.Contains(result, "WAR!") {
		t.Errorf("Expected result to contain 'WAR!', got: %s", result)
	}

	// Test cards at stake (2 initial + 8 war cards = 10 total)
	if warInfo.CardsAtStake != 10 {
		t.Errorf("Expected 10 cards at stake, got %d", warInfo.CardsAtStake)
	}

	// Test war count
	if warInfo.WarCount != 1 {
		t.Errorf("Expected 1 war, got %d", warInfo.WarCount)
	}

	// Player should win with 10 > 5
	if !strings.Contains(result, "Player wins!") {
		t.Errorf("Expected player to win, got: %s", result)
	}
}
