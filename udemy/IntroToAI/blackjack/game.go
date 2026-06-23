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

	// Initial deal: two cards per PlayeRound
	for range 2 {
		human.AddCard(deck.Draw(), cardCounter)
		ai.AddCard(deck.Draw(), cardCounter)
		dealer.AddCard(deck.Draw(), cardCounter)
	}

	// Show initial hands
	fmt.Println("\nIntial Deal:")
	dealer.DisplayHand(true) // Hiding dealer's second card
	human.DisplayHand(false)
	ai.DisplayHand(false)

	// Play each player's turn
	human.PlayTurn(deck, cardCounter, dealer.Hand[0]) // Show dealer's up card
	if !human.IsBust {
		// let ai player play
		ai.PlayTurn(deck, cardCounter, dealer.Hand[0]) // Show dealer's up card
		// let dealer play
	}

	// Show results

	// Display results

	// Display card counting statistics
}
