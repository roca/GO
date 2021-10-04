package main

/*
1. Reshuffling
2. add betting
3. blackjack payouts (1.5x)
4. doubling down
5. splitting 7,7
*/

import (
	"deck_of_cards/blackjack"
)

func main() {
	game := blackjack.New()
	_ = game.Play(blackjack.HumanAI())
}
