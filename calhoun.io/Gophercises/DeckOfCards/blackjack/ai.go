package blackjack

import (
	"deck_of_cards/deck"
	"fmt"
)

func HumanAI() AI {
	return humanAI{}
}

type humanAI struct{}

func (ai humanAI) Bet(shuffled bool) int {
	if shuffled {
		fmt.Println("The deck was just shuffled!")
	}
	fmt.Println("What would you like to bet?")
	var bet int
	fmt.Scanf("%d\n", &bet)
	return bet
}

func (ai humanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
		fmt.Println("Player:", hand, Score(hand...))
		fmt.Println("What will you do? (h)it, (s)and, (d)ouble")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		case "d":
			return MoveDouble
		default:
			fmt.Println("Invalid option:", input)
		}
	}
}

func (ai humanAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:")
	for i, h := range hands {
		fmt.Println(" ", i, ":", h, Score(h...))
	}
	fmt.Println("Dealer:", dealer, Score(dealer...))
	fmt.Println("===============")
}
