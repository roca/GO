package test

import (
	"os"
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
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of 'Four of Clubs', bot got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := deck.NewDeck()
	d.SaveToFile("_decktesting")

	loadedDeck := deck.NewDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}
}
