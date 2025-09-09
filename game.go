package main

import "fmt"

func StartGame() (*Player, *Player) {
	// 1. Create two players
	player1 := Player{Name: "Player 1", Cards: []Card{}}
	cpu := Player{Name: "CPU", Cards: []Card{}}

	// 2. Create a deck
	deck := NewDeck()
	fmt.Printf("Deck created with %d cards\n", len(deck.Cards))

	// 3. Shuffle the deck
	deck.Shuffle()
	fmt.Println("Deck shuffled!")

	// 4. Deal cards to both players (28 each)
	for i := range 28 {
		player1.AddCard(deck.Cards[i])
		cpu.AddCard(deck.Cards[i+28])
	}
	return &player1, &cpu // Return pointers to the players

}

func PlayRound(player1 *Player, cpu *Player) string {
	// Check if either player can't war, this is a game over condition
	if !player1.HasCards() || !cpu.HasCards() {
		return "Game over - someone ran out of cards"
	}

	allCards := []Card{} // All cards that will go to the winner

	// Initial cards
	card1 := player1.PlayCard()
	card2 := cpu.PlayCard()
	allCards = append(allCards, card1, card2)

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

	return result
}

// Helper function to handle war card placement
func putDownWarCards(player *Player) ([]Card, Card) {
	warCards := []Card{}
	var lastCard Card

	// Put down up to 4 cards, or whatever the player has
	cardsToPlay := min(4, len(player.Cards))

	for i := 0; i < cardsToPlay; i++ {
		card := player.PlayCard()
		warCards = append(warCards, card)
		lastCard = card // Keep track of the last one
	}

	return warCards, lastCard
}
