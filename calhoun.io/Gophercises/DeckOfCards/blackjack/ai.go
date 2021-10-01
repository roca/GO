package blackjack

import (
	"deck_of_cards/deck"
	"fmt"
)

func HumanAI() AI {
	return humanAI{}
}

type humanAI struct{}

func (ai humanAI) Bet() int {
	return 1
}

func (ai humanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
		fmt.Println("Player:", hand, Score(hand...))
		fmt.Println("What will you do? (h)it, (s)and")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		default:
			fmt.Println("Invalid option:", input)
		}
	}
}

func (ai humanAI) Results(hand [][]deck.Card, dealer []deck.Card) {
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:", hand, Score(hand[0]...))
	fmt.Println("Dealer:", dealer, Score(dealer...))
	fmt.Println("===============")
}
