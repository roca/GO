package main

/*
1. Reshuffling
2. add betting
3. blackjack payouts (1.5x)
4. doubling down
5. TODO: splitting 7,7
*/

import (
	"deck_of_cards/blackjack"
	"fmt"
)

func main() {
	startDefaults := blackjack.StartOption{
		Decks:           4,
		Hands:           50000,
		BlackjackPayout: 1.5,
	}
	game := blackjack.New(startDefaults)
	winnings := game.Play(blackjack.BasicAI())

	fmt.Println(winnings)
}
