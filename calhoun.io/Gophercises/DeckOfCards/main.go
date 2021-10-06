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
)

func main() {
	startDefaults := blackjack.StartOption{
		Decks:            3,
		Hands:            2,
		BlackjackPayout: 1.5,
	}
	game := blackjack.New(startDefaults)
	_ = game.Play(blackjack.HumanAI())
}
