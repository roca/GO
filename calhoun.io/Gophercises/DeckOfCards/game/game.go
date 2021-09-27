package game

import (
	"deck_of_cards/deck"
	"fmt"
)

type State int8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

type GameState struct {
	Deck   []deck.Card
	State  State
	Player deck.Hand
	Dealer deck.Hand
}

func (gs *GameState) CurrentPlayer() *deck.Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("it isn't currently any player's turn")
	}
}

func Shuffle(gs GameState) GameState {
	ret := clone(gs)
	ret.Deck = deck.Shuffle(ret.Deck)
	return ret
}

func Deal(gs GameState) GameState {
	ret := clone(gs)
	ret.Player = make(deck.Hand, 0, 5)
	ret.Dealer = make(deck.Hand, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		for _, hand := range []*deck.Hand{&ret.Player, &ret.Dealer} {
			card, ret.Deck = deck.Draw(ret.Deck)
			*hand = append(*hand, card)
		}
	}
	ret.State = StatePlayerTurn
	return ret
}

func Hit(gs GameState) GameState {
	ret := clone(gs)
	hand := ret.CurrentPlayer()
	var card deck.Card
	card, ret.Deck = deck.Draw(ret.Deck)
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		return Stand(ret)
	}
	return ret
}

func Stand(gs GameState) GameState {
	ret := clone(gs)
	ret.State++
	return ret
}

func EndHand(gs GameState) GameState {
	ret := clone(gs)
	pScore, dScore := ret.Player.Score(), ret.Dealer.Score()
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:", ret.Player, ",Score:", pScore)
	fmt.Println("Dealer:", ret.Dealer, ",Score:", dScore)
	switch {
	case pScore > 21:
		fmt.Println("You busted")
	case dScore > 21:
		fmt.Println("Dealer busted")
	case pScore > dScore:
		fmt.Println("You win!")
	case dScore > pScore:
		fmt.Println("You loose")
	case dScore == pScore:
		fmt.Println("Draw")
	}
	fmt.Printf("\n %d cards left in deck \n\n", len(ret.Deck))

	ret.Player = nil
	ret.Dealer = nil

	return ret
}

func clone(gs GameState) GameState {
	ret := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make(deck.Hand, len(gs.Player)),
		Dealer: make(deck.Hand, len(gs.Dealer)),
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)
	return ret
}
