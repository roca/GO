package main

import "fmt"

const (
	Hearts   = "\u2665"
	Diamonds = "\u2666"
	Clubs    = "\u2663"
	Spades   = "\u2660"
)

type Card struct {
	Suit  string
	Value string
	Score int
}

func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Value, c.Suit)
}
