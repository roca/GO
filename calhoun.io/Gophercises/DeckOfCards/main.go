package main

import (
	"deck_of_cards/blackjack"
	"deck_of_cards/deck"
	"deck_of_cards/game"
)

func main() {
	// cards := deck.New(deck.Deck(3), deck.Shuffle)
	var gs game.GameState
	gs.Deck = deck.New(deck.Deck(3))
	gs = game.Shuffle(gs)

	for i := 0; i < 10; i++ {
		gs = game.Deal(gs)

		player := blackjack.NewPlayer(gs)

		for player.GameState.State == game.StatePlayerTurn {
			player.GameState = player.Play()(player.GameState)
		}

		dealer := blackjack.NewDealer(gs)

		for dealer.GameState.State == game.StateDealerTurn {
			dealer.GameState = dealer.Play()(dealer.GameState)
		}

		blackjack.Results([]game.GameState{gs})
	}
}
