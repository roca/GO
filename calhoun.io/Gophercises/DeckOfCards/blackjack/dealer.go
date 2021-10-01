package blackjack

import (
	"deck_of_cards/deck"
)

func DealerAI() AI {
	return dealerAI{}
}

type dealerAI struct{}

func (ai dealerAI) Bet() int {
	// noop
	return 1
}

func (ai dealerAI) Play(hand []deck.Card, dealer deck.Card) Move {
	dScore := Score(hand...)
	if dScore <= 16 || (dScore == 17 && Soft(hand...)) {
		return MoveHit
	}
	return MoveStand
}

func (ai dealerAI) Results(hand [][]deck.Card, dealer []deck.Card) {
	// noop
}
