package blackjack

import (
	"deck_of_cards/deck"
	"deck_of_cards/game"
)

type DealerAI struct {
	GameState game.GameState
}

func NewDealer(gs game.GameState) DealerAI {
	return DealerAI{
		GameState: gs,
	}
}

func (ai *DealerAI) Bet() int {
	return 1
}

func (ai *DealerAI) Play() Move {
	// If dealer score <= 16, we hit
	// If dealer has a soft 17, then we hit.
	if ai.GameState.Dealer.Score() <= 16 || (ai.GameState.Dealer.Score() == 17 && ai.GameState.Dealer.MinScore() != 17) {
		return Hit
	} else {
		return Stand
	}
}

func (ai *DealerAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	for _, hand := range hands {
		gs := game.GameState{
			Player: hand,
			Dealer: dealer,
		}
		gs = game.EndHand(gs)
	}
}
