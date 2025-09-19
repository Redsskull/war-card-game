package main

import (
	"fmt"
	"math/rand"
)

func StartGame() (*Player, *Player) {
	// 1. Create two players
	player1 := Player{Name: "Player 1", Cards: []Card{}}
	cpu := Player{Name: "CPU", Cards: []Card{}}

	// 2. Create a deck
	deck := NewDeck()

	// 3. Shuffle the deck
	deck.Shuffle()

	// 4. Deal cards evenly to both players, give leftover to random player
	totalCards := len(deck.Cards)
	cardsPerPlayer := totalCards / 2 // 55/2 = 27 (integer division)

	// Deal base cards to each player
	for i := range cardsPerPlayer {
		player1.AddCard(deck.Cards[i])            // Player gets 0-26
		cpu.AddCard(deck.Cards[i+cardsPerPlayer]) // CPU gets 27-53
	}

	// Deal leftover card(s) randomly
	leftoverCards := totalCards - (cardsPerPlayer * 2) // 55 - 54 = 1 leftover
	for i := range leftoverCards {
		cardIndex := cardsPerPlayer*2 + i // Index 54 for the leftover card
		if rand.Intn(2) == 0 {            // Random 0 or 1
			player1.AddCard(deck.Cards[cardIndex])
		} else {
			cpu.AddCard(deck.Cards[cardIndex])
		}
	}

	return &player1, &cpu
}

func PlayRound(player1 *Player, cpu *Player) (Card, Card, string) {
	// Check if either player can't war, this is a game over condition
	if !player1.HasCards() || !cpu.HasCards() {
		return Card{}, Card{}, "Game over - someone ran out of cards"
	}

	allCards := []Card{} // All cards that will go to the winner

	// Initial cards
	card1 := player1.PlayCard()
	card2 := cpu.PlayCard()
	allCards = append(allCards, card1, card2)

	// Keep track of the final cards for display
	finalCard1 := card1
	finalCard2 := card2

	result := fmt.Sprintf("Player: %s, CPU: %s", card1, card2)

	// Handle wars (only if there's a tie)
	for card1.Value == card2.Value {
		result += " -> WAR!"

		// Each player puts down cards for war
		warCards1, lastCard1 := putDownWarCards(player1)
		warCards2, lastCard2 := putDownWarCards(cpu)

		allCards = append(allCards, warCards1...)
		allCards = append(allCards, warCards2...)

		card1, card2 = lastCard1, lastCard2
		finalCard1, finalCard2 = card1, card2 // Update final cards
		result += fmt.Sprintf(" -> Player: %s, CPU: %s", card1, card2)
	}

	// Now determine winner (works for both simple comparison AND after wars)
	if card1.Value > card2.Value {
		// Player 1 wins
		for _, card := range allCards {
			player1.AddCard(card)
		}
		result += " -> Player wins!"
	} else {
		// CPU wins
		for _, card := range allCards {
			cpu.AddCard(card)
		}
		result += " -> CPU wins!"
	}

	return finalCard1, finalCard2, result
}

// Helper function to handle war card placement
func putDownWarCards(player *Player) ([]Card, Card) {
	warCards := []Card{}
	lastCard := Card{}

	// Put down up to 4 cards, or whatever the player has
	cardsToPlay := min(4, len(player.Cards))

	for range cardsToPlay {
		card := player.PlayCard()
		warCards = append(warCards, card)
		lastCard = card
	}
	return warCards, lastCard
}

// Check if the game is over and return winner info
func IsGameOver(player1, cpu *Player) (bool, string) {
	if !player1.HasCards() {
		return true, "GAME OVER! CPU WINS! 🤖"
	} else if !cpu.HasCards() {
		return true, "GAME OVER! YOU WIN! 🏆"
	}
	return false, ""
}

// ExecuteGameRound plays one complete round and returns the cards played,
// result message, and whether the game is over with winner info.
func ExecuteGameRound(player1, cpu *Player) (Card, Card, string, bool, string) {
	// Check if game can continue
	if !player1.HasCards() || !cpu.HasCards() {
		gameOver, winner := IsGameOver(player1, cpu)
		return Card{}, Card{}, "", gameOver, winner
	}

	// Play the round
	playerCard, cpuCard, result := PlayRound(player1, cpu)

	// Check if game is now over
	gameOver, winner := IsGameOver(player1, cpu)

	return playerCard, cpuCard, result, gameOver, winner
}
