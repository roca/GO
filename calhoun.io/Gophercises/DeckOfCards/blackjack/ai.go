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

func NewPlayer(gs game.GameState) HumanAI {
	return HumanAI{
		GameState: gs,
	}
}

func (ai *HumanAI) Bet() int {
	return 1
}

func (ai *HumanAI) Play() Move {
	for {
		fmt.Println("Player:", ai.GameState.Player, ",Score:", ai.GameState.Player.Score())
		fmt.Println("Dealer:", ai.GameState.Dealer[0])
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

func Results(games []game.GameState) {
	for _, gs := range games {
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
