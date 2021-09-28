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

	//for i := 0; i < 10; i++ {
		gs = game.Deal(gs)

		player := blackjack.New(gs)

		for player.GameState.State == game.StatePlayerTurn {
			player.GameState = player.Play(player.GameState.Player, player.GameState.Dealer[0])(player.GameState)
		}

		for player.GameState.State == game.StateDealerTurn {
			// If dealer score <= 16, we hit
			// If dealer has a soft 17, then we hit.
			if player.GameState.Dealer.Score() <= 16 || (player.GameState.Dealer.Score() == 17 && player.GameState.Dealer.MinScore() != 17) {
				player.GameState = game.Hit(player.GameState)
			} else {
				player.GameState = game.Stand(player.GameState)
			}
		}

		player.Results([][]deck.Card{player.GameState.Player}, player.GameState.Dealer)
	//}
}
