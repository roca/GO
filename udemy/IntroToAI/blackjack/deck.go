package main

import (
	"fmt"
	"math/rand"
)

type Deck []Card

func NewDeck() Deck {
	deck := Deck{}
	suits := []string{Hearts, Diamonds, Clubs, Spades}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	scores := []int{11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}

	// Create a standard deck of 52 cards
	for _, suit := range suits {
		for i, value := range values {
			card := Card{
				Suit:  suit,
				Value: value,
				Score: scores[i],
			}
			deck = append(deck, card)
		}
	}

	return deck
}

func (d Deck) Shuffle() Deck {
	// Create a shuffled deck of cards
	shuffled := make(Deck, len(d))
	copy(shuffled, d)

	// Fisher-Yates shuffle algorithm
	for i := len(shuffled) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	return shuffled
}

func (d *Deck) Draw() Card {
	// If Deck is empty...
	if len(*d) == 0 {
		fmt.Println("Deck is empty! Reshuffling...")
		*d = NewDeck().Shuffle()
	}

	card := (*d)[0]
	return card
}
