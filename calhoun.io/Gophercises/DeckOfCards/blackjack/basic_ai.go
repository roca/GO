package blackjack

import (
	"deck_of_cards/deck"
)

type basicAI struct {
	score int
	seen  int
	decks int
}

type BasicAIOption struct {
	Score int
	Seen  int
	Decks int
}

func BasicAI(opts ...interface{}) AI {
	aiDefaults := BasicAIOption{}
	for _, opt := range opts {
		switch o := opt.(type) {
		case BasicAIOption:
			aiDefaults = o
		}
	}

	return basicAI{
		score: aiDefaults.Score,
		seen:  aiDefaults.Seen,
		decks: aiDefaults.Decks,
	}
}

func (ai basicAI) Bet(shuffled bool) int {
	refAI := &ai
	if shuffled {
		refAI.score = 0
		refAI.seen = 0
	}
	trueScore := ai.score / ((ai.decks*52 - ai.seen) / 52)
	switch {
	case trueScore >= 14:
		return 100000
	case trueScore >= 8:
		return 5000
	default:
		return 100
	}
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
	for _, card := range dealer {
		ai.count(card)
	}
	for _, hand := range hands {
		for _, card := range hand {
			ai.count(card)
		}
	}
}
func (ai basicAI) count(card deck.Card) {
	refAI := &ai
	score := Score(card)
	switch {
	case score >= 10:
		refAI.score--
	case score <= 6:
		refAI.score++
	}
	refAI.seen++
}
