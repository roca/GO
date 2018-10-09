package main

import "github.com/GOCODE/udemy/gocast/cards/deck"

func main() {
	cards := deck.NewDeck()
	cards.Shuffle()
	cards.Print()
}
