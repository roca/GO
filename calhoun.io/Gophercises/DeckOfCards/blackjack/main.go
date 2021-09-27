package main

import (
	"deck_of_cards/deck"
	"deck_of_cards/game"
	"fmt"
)

func main() {
	// cards := deck.New(deck.Deck(3), deck.Shuffle)
	var gs game.GameState
	gs.Deck = deck.New(deck.Deck(3))
	gs = game.Shuffle(gs)

	for i := 0; i < 10; i++ {
		gs = game.Deal(gs)

		var input string
		for gs.State == game.StatePlayerTurn {
			fmt.Println("Player:", gs.Player, ",Score:", gs.Player.Score())
			fmt.Println("Dealer:", gs.Dealer.DealerString())
			fmt.Println("What will you do? (h)it or (s)tand")
			fmt.Scanf("%s\n", &input)
			switch input {
			case "h":
				gs = game.Hit(gs)
			case "s":
				gs = game.Stand(gs)
			default:
				fmt.Println("Invalid option:", input)
			}
		}

		for gs.State == game.StateDealerTurn {
			// If dealer score <= 16, we hit
			// If dealer has a soft 17, then we hit.
			if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
				gs = game.Hit(gs)
			} else {
				gs = game.Stand(gs)
			}
		}

		gs = game.EndHand(gs)
	}
}
