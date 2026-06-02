package main

import "fmt"

func PlayeRound(deck *Deck, cardCounter *CardCounter) {
	fmt.Println("\n=== New Round of BlackJack ===")
	fmt.Printf("Cards remaining in deck: %d\n", len(*deck))

	// Initialize players
	dealer := NewPlayer("Dealer", false)
	human := NewPlayer("Human", false)
	ai := NewPlayer("AI", true)

	fmt.Println("Players:", dealer.Name, human.Name, ai.Name)

	// Initial deal: two cards per PlayeRound(

	// Show initial hands

	// Play each player's turn

	// Show results

	// Display results

	// Display card counting statistics
}
