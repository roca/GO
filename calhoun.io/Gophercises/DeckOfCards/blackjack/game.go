package blackjack

import (
	"deck_of_cards/deck"
	"fmt"
)

const (
	statePlayerTurn state = iota
	stateDealerTurn
	stateHandOver
)

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

type state int8

type Game struct {
	// unexported fields
	nDecks          int
	nHands          int
	blackjackPayout float64

	state state
	deck  []deck.Card

	player    []deck.Card
	playerBet int
	balance   int

	dealer   []deck.Card
	dealerAI AI
}

type AI interface {
	Bet(shuffled bool) int
	Play(hand []deck.Card, dealer deck.Card) Move
	Results(hands [][]deck.Card, dealer []deck.Card)
}

type StartOption struct {
	Decks           int
	Hands           int
	BlackjackPayout float64
}

func (o StartOption) String() string {
	return fmt.Sprintf("{Decks: %d, Hands: %d, BlackjackPayout: %g}", o.Decks, o.Hands, o.BlackjackPayout)
}

func New(opts ...interface{}) Game {
	startDefaults := StartOption{3, 100, 1.5}
	for _, opt := range opts {
		switch o := opt.(type) {
		case StartOption:
			startDefaults = o
		}
	}
	fmt.Println("New Game started with default options", startDefaults)
	return Game{
		nDecks:          startDefaults.Decks,
		nHands:          startDefaults.Hands,
		state:           statePlayerTurn,
		dealerAI:        dealerAI{},
		balance:         0,
		blackjackPayout: startDefaults.BlackjackPayout,
	}
}

func (g *Game) currentHand() *[]deck.Card {
	switch g.state {
	case statePlayerTurn:
		return &g.player
	case stateDealerTurn:
		return &g.dealer
	default:
		panic("it isn't currently any player's turn")
	}
}

func bet(g *Game, ai AI, shuffled bool) {
	g.playerBet = ai.Bet(shuffled)
}

func deal(g *Game) {
	g.player = make([]deck.Card, 0, 5)
	g.dealer = make([]deck.Card, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, g.deck = draw(g.deck)
		g.player = append(g.player, card)
		card, g.deck = draw(g.deck)
		g.dealer = append(g.dealer, card)
	}
	g.state = statePlayerTurn
}

func (g *Game) Play(ai AI) int {
	g.deck = nil
	minCards := 52 * g.nDecks / 3
	for i := 0; i < g.nHands; i++ {
		shuffled := false
		if len(g.deck) < minCards {
			g.deck = deck.New(deck.Deck(g.nDecks), deck.Shuffle)
			shuffled = true
		}
		bet(g, ai, shuffled)
		deal(g)
		if Blackjack(g.dealer...) {
			endHand(g, ai)
			continue
		}
		for g.state == statePlayerTurn {
			hand := make([]deck.Card, len(g.player))
			copy(hand, g.player)
			move := ai.Play(hand, g.dealer[0])
			move(g)
		}
		for g.state == stateDealerTurn {
			hand := make([]deck.Card, len(g.dealer))
			copy(hand, g.dealer)
			move := g.dealerAI.Play(hand, g.dealer[0])
			move(g)
		}
		endHand(g, ai)
	}
	return g.balance
}

type Move func(*Game)

func MoveHit(g *Game) {
	hand := g.currentHand()
	var card deck.Card
	card, g.deck = draw(g.deck)
	*hand = append(*hand, card)
	if Score(*hand...) > 21 {
		MoveStand(g)
	}
}

func MoveStand(g *Game) {
	g.state++
}

func endHand(g *Game, ai AI) {
	pScore, dScore := Score(g.player...), Score(g.dealer...)
	// TODO(roca): Figure out winnings and add/subtract them
	pBlackjack, dBlackjack := Blackjack(g.player...), Blackjack(g.dealer...)
	winnings := g.playerBet
	switch {
	case pBlackjack && dBlackjack:
		winnings = 0
	case dBlackjack:
		winnings = -winnings
	case pBlackjack:
		winnings = int(float64(winnings) * g.blackjackPayout)
	case pScore > 21:
		winnings = -winnings
	case dScore > 21:
		// win
	case pScore > dScore:
		// win
	case dScore > pScore:
		winnings = -winnings
	case dScore == pScore:
		winnings = 0
	}
	g.balance += winnings
	ai.Results([][]deck.Card{g.player}, g.dealer)
	fmt.Printf("Balance: %d\n", g.balance)
	g.player = nil
	g.dealer = nil
}

// Score will take in a hand of cards and return the best blackjack score
// possible with the hjand.
func Score(hand ...deck.Card) int {
	minScore := minScore(hand...)
	if minScore > 11 {
		return minScore
	}
	for _, card := range hand {
		// Ace is currently worth 1, and we are changing it to be worth 11
		// 11 - 1 = 10
		if card.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

// Soft returns true if the score of a hand is a soft score - that is if an ace
// is beeing counted as 11 points
func Soft(hand ...deck.Card) bool {
	minScore := minScore(hand...)
	score := Score(hand...)
	return minScore != score
}

// Blackjack returns true if a hand is a blackjack
func Blackjack(hand ...deck.Card) bool {
	return len(hand) == 2 && Score(hand...) == 21
}

func minScore(hand ...deck.Card) int {
	var score int
	for _, c := range hand {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
