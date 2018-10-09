package test

import (
	"testing"

	"github.com/GOCODE/udemy/gocast/cards/deck"
)

func TestNewDeck(t *testing.T) {
	d := deck.NewDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of 'Ace of Spades', but got %v", d[0])
	}
}
