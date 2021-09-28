package blackjack

import (
	"deck_of_cards/deck"
	"deck_of_cards/game"
	"fmt"
)

type AI interface {
	Bet() int
	Play(hand []deck.Card, dealer deck.Card) Move
	Results(hands [][]deck.Card, dealer []deck.Card)
}

type HumanAI struct {
	GameState game.GameState
}

func New(gs game.GameState) HumanAI {
	return HumanAI{
		GameState: gs,
	}
}

func (ai *HumanAI) Bet() int {
	return 1
}

func (ai *HumanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
		fmt.Println("Player:", hand, ",Score:", deck.Hand(hand).Score())
		fmt.Println("Dealer:", dealer)
		fmt.Println("What will you do? (h)it or (s)tand")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return Hit
		case "s":
			return Stand
		default:
			fmt.Println("Invalid option:", input)
		}
	}
}

type Move func(game.GameState) game.GameState

func (ai *HumanAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	for _, hand := range hands {
		gs := game.GameState{
			Player: hand,
			Dealer: dealer,
		}
		gs = game.EndHand(gs)
	}
}

// Filler to be implemented later
func Hit(gs game.GameState) game.GameState {
	return game.Hit(gs)
}

func Stand(gs game.GameState) game.GameState {
	return game.Stand(gs)
}
