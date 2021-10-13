package blackjack

import "deck_of_cards/deck"

type basicAI struct{}

func BasicAI() AI {
	return basicAI{}
}

func (ai basicAI) Bet(shuffled bool) int {
	return 100
}

func (ai basicAI) Play(hand []deck.Card, dealer deck.Card) Move {
	score := Score(hand...)
	if len(hand) == 2 {
		if hand[0] == hand[1] {
			cardScore := Score(hand[0])
			if cardScore >= 8 && cardScore != 10 {
				return MoveSplit
			}

		}
		if (score == 10 || score == 11) && !Soft(hand...) {
			return MoveDouble
		}
	}
	dScore := Score(dealer)
	if dScore >= 5 && dScore <= 6 {
		return MoveStand
	}
	if score < 13 {
		return MoveHit
	}
	return MoveStand
}

func (ai basicAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	// noop
}
