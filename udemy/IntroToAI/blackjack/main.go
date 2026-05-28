package main

import "fmt"

func main() {
	// Clear the screen
	clearScreen()

	// Show a welcome message
	fmt.Println("=== Go Blackjack ===")
	fmt.Println("Welcome to Go Blackjack!\nTry to get as close to 21 as possible without going over.")
	fmt.Println("You're against an AI opponent with card counting abilities. Good luck!")

	// Get a deck of cards
	deck := NewDeck().Shuffle()

	for _, c := range deck {
		fmt.Printf("%s%s\n", c.Value, c.Suit)
	}

	// Create a card counter for the AI to use

	for {
		// Check to see if we need to shuffle the deck

		// Play a round

		// Ask if the player wants to play another round

		// If not, quit the game

		// clear the screen
	}
}
