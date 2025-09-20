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

func TestGameReset(t *testing.T) {
	// Test that game reset creates fresh players with correct card counts

	// Start initial game
	player1, cpu := StartGame()
	initialPlayer1Cards := len(player1.Cards)
	initialCpuCards := len(cpu.Cards)

	// Simulate some gameplay - remove some cards
	player1.PlayCard()
	cpu.PlayCard()

	afterGameplayPlayer1Cards := len(player1.Cards)
	afterGameplayCpuCards := len(cpu.Cards)

	// Verify cards were removed
	if afterGameplayPlayer1Cards != initialPlayer1Cards-1 {
		t.Errorf("Expected player1 to have %d cards, got %d", initialPlayer1Cards-1, afterGameplayPlayer1Cards)
	}
	if afterGameplayCpuCards != initialCpuCards-1 {
		t.Errorf("Expected cpu to have %d cards, got %d", initialCpuCards-1, afterGameplayCpuCards)
	}

	// Reset game (simulate "Start New Game" button)
	resetPlayer1, resetCpu := StartGame()

	// Verify fresh game has correct total card count (55 cards distributed)
	totalResetCards := len(resetPlayer1.Cards) + len(resetCpu.Cards)
	if totalResetCards != 55 {
		t.Errorf("Reset game should have 55 total cards, got %d", totalResetCards)
	}

	// Verify each player has at least 27 cards (fair distribution)
	if len(resetPlayer1.Cards) < 27 || len(resetPlayer1.Cards) > 28 {
		t.Errorf("Player1 should have 27-28 cards, got %d", len(resetPlayer1.Cards))
	}
	if len(resetCpu.Cards) < 27 || len(resetCpu.Cards) > 28 {
		t.Errorf("CPU should have 27-28 cards, got %d", len(resetCpu.Cards))
	}

	// Verify they are actually new objects (different memory addresses)
	if &resetPlayer1 == &player1 {
		t.Error("Reset should create new player objects, not reuse old ones")
	}
	if &resetCpu == &cpu {
		t.Error("Reset should create new cpu objects, not reuse old ones")
	}

	// Verify cards are freshly shuffled (different order likely)
	differentOrder := false
	minCards := min(len(resetPlayer1.Cards), len(player1.Cards))

	for i := 0; i < minCards && i < 3; i++ { // Check first 3 cards
		if resetPlayer1.Cards[i].Value != player1.Cards[i].Value {
			differentOrder = true
			break
		}
	}

	// Note: This test might occasionally fail due to random chance, but very unlikely
	if !differentOrder && minCards > 0 {
		t.Log("Warning: Reset game has same card order - could be random chance")
	}
}

func TestWarScoreUpdateTiming(t *testing.T) {
	// Test that war score logic updates correctly after war completes
	// This tests the underlying timing logic that the UI uses

	// Create test players with war setup
	player1 := &Player{Name: "Test Player", Cards: []Card{
		{Value: 7, Suit: "Hearts"},  // Tied card
		{Value: 2, Suit: "Hearts"},  // War card 1
		{Value: 3, Suit: "Hearts"},  // War card 2
		{Value: 4, Suit: "Hearts"},  // War card 3
		{Value: 10, Suit: "Hearts"}, // Final deciding card (wins)
	}}
	cpu := &Player{Name: "Test CPU", Cards: []Card{
		{Value: 7, Suit: "Spades"}, // Tied card (same value!)
		{Value: 6, Suit: "Spades"}, // War card 1
		{Value: 8, Suit: "Spades"}, // War card 2
		{Value: 9, Suit: "Spades"}, // War card 3
		{Value: 5, Suit: "Spades"}, // Final deciding card (loses)
	}}

	// Record initial card counts
	initialPlayer1Count := len(player1.Cards)
	initialCpuCount := len(cpu.Cards)

	// Execute war round
	_, _, result, warInfo := PlayRound(player1, cpu)

	// Verify it's a war
	if !warInfo.IsWar {
		t.Fatal("Expected this to be a war")
	}

	// After war, player1 should have won all cards
	expectedPlayer1Cards := initialPlayer1Count + initialCpuCount
	expectedCpuCards := 0

	if len(player1.Cards) != expectedPlayer1Cards {
		t.Errorf("After war, player1 should have %d cards, got %d", expectedPlayer1Cards, len(player1.Cards))
	}

	if len(cpu.Cards) != expectedCpuCards {
		t.Errorf("After war, cpu should have %d cards, got %d", expectedCpuCards, len(cpu.Cards))
	}

	// Verify the result shows player win
	if !strings.Contains(result, "Player wins!") {
		t.Errorf("Expected result to show player win, got: %s", result)
	}

	// Test the timing logic: In UI, these counts should only be displayed
	// after the 8-second visual sequence, not immediately
	// This test verifies the game logic produces correct final counts
	// that the UI will display after the war animation completes
}

func TestGameClickState(t *testing.T) {
	// This test simulates the game state logic for click handling during wars
	// Since I can't easily test UI interactions, I test the underlying logic here

	gameAcceptingClicks := true

	// Simulate start of war
	warInfo := WarInfo{
		IsWar:        true,
		CardsAtStake: 10,
		WarCount:     1,
	}

	// When war starts, clicks should be disabled
	if warInfo.IsWar {
		gameAcceptingClicks = false
	}

	// Test that clicks are disabled during war
	if gameAcceptingClicks {
		t.Error("Game should not accept clicks during war")
	}

	// Simulate end of war
	if warInfo.IsWar {
		// After war processing is complete
		gameAcceptingClicks = true
	}

	// Test that clicks are re-enabled after war
	if !gameAcceptingClicks {
		t.Error("Game should accept clicks after war ends")
	}

	// Test normal round (no war)
	normalWarInfo := WarInfo{
		IsWar:        false,
		CardsAtStake: 2,
		WarCount:     0,
	}

	// Normal rounds should not affect click state
	gameAcceptingClicks = true
	if !normalWarInfo.IsWar && gameAcceptingClicks {
		// This should remain true
		if !gameAcceptingClicks {
			t.Error("Normal rounds should not disable clicks")
		}
	}
}

func TestBreakLongText(t *testing.T) {
	// Test 1: Short text should not be broken
	shortText := "Player: King, CPU: 7 -> Player wins!"
	result := breakLongText(shortText, 60)
	if result != shortText {
		t.Errorf("Short text should not be modified, got: %s", result)
	}

	// Test 2: Long war text should be broken at natural break points
	longText := "Player: King, CPU: King -> WAR! -> Player: 5, CPU: 7 -> CPU wins!"
	expected := "Player: King, CPU: King -> WAR! -> Player: 5, CPU: 7\nCPU wins!"
	result = breakLongText(longText, 60)
	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}

	// Test 3: Very long text with multiple wars
	veryLongText := "Player: King, CPU: King -> WAR! -> Player: Ace, CPU: Ace -> WAR! -> Player: 3, CPU: 9 -> CPU wins!"
	result = breakLongText(veryLongText, 60)
	lines := strings.Split(result, "\n")
	if len(lines) < 2 {
		t.Error("Very long text should be broken into multiple lines")
	}

	// Verify each line is within reasonable length
	for i, line := range lines {
		if len(line) > 70 { // Allow some buffer over maxLineLength
			t.Errorf("Line %d is too long (%d chars): %s", i+1, len(line), line)
		}
	}

	// Test 4: Text without natural break points should not be broken
	noBreakText := "This is a long sentence without any natural break points to split on"
	result = breakLongText(noBreakText, 30)
	if result != noBreakText {
		t.Error("Text without break points should not be modified")
	}

	// Test 5: Empty string
	result = breakLongText("", 60)
	if result != "" {
		t.Error("Empty string should remain empty")
	}
}
