package main

import (
	"deck_of_cards/blackjack"
)

func main() {
	game := blackjack.New()
	_ = game.Play(blackjack.HumanAI())
}
