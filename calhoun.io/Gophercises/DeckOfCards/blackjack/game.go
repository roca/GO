package blackjack

import (
	"deck_of_cards/deck"
	"errors"
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

type hand struct {
	cards []deck.Card
	bet   int
}

type Game struct {
	// unexported fields
	nDecks          int
	nHands          int
	blackjackPayout float64

	state state
	deck  []deck.Card

	player    []hand
	handIdx   int
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
	// fmt.Println("New Game started with default options", startDefaults)
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
		return &g.player[g.handIdx].cards
	case stateDealerTurn:
		return &g.dealer
	default:
		panic("it isn't currently any player's turn")
	}
}

func bet(g *Game, ai AI, shuffled bool) {
	bet := ai.Bet(shuffled)
	if bet < 100 {
		panic("bet must be at least 100")
	}
	g.playerBet = bet
}

func deal(g *Game) {
	playerHand := make([]deck.Card, 0, 5)
	g.handIdx = 0
	g.dealer = make([]deck.Card, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, g.deck = draw(g.deck)
		playerHand = append(playerHand, card)
		card, g.deck = draw(g.deck)
		g.dealer = append(g.dealer, card)
	}
	g.player = []hand{
		{cards: playerHand, bet: g.playerBet},
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
			endRound(g, ai)
			continue
		}
		for g.state == statePlayerTurn {
			hand := make([]deck.Card, len(*g.currentHand()))
			copy(hand, *g.currentHand())
			move := ai.Play(hand, g.dealer[0])
			err := move(g)
			switch err {
			case errBust:
				_ = MoveStand(g)
			case nil:
				// noop
			default:
				panic(err)
			}
		}
		for g.state == stateDealerTurn {
			hand := make([]deck.Card, len(g.dealer))
			copy(hand, g.dealer)
			move := g.dealerAI.Play(hand, g.dealer[0])
			move(g)
		}
		endRound(g, ai)
	}
	return g.balance
}

var (
	errBust = errors.New("hand score exceed 21")
)

type Move func(*Game) error

func MoveHit(g *Game) error {
	hand := g.currentHand()
	var card deck.Card
	card, g.deck = draw(g.deck)
	*hand = append(*hand, card)
	if Score(*hand...) > 21 {
		return errBust
	}
	return nil
}

func MoveSplit(g *Game) error {
	cards := g.currentHand()
	if len(*cards) != 2 {
		return errors.New("you can only split with two cards in your hand")
	}
	if (*cards)[0].Rank != (*cards)[1].Rank {
		return errors.New("both cards must have the same rank to split")
	}
	g.player = append(g.player, hand{
		cards: []deck.Card{(*cards)[1]},
		bet:   g.player[g.handIdx].bet,
	})
	g.player[g.handIdx].cards = (*cards)[:1]
	return nil
}

func MoveDouble(g *Game) error {
	if len(*g.currentHand()) != 2 {
		return errors.New("can only double on a hand with 2 cards")
	}
	g.playerBet *= 2
	_ = MoveHit(g)
	return MoveStand(g)
}

func MoveStand(g *Game) error {

	switch g.state {
	case stateDealerTurn:
		g.state++
		return nil
	case statePlayerTurn:
		g.handIdx++
		if g.handIdx >= len(g.player) {
			g.state++
		}
		return nil
	default:
		return errors.New("invalid state")
	}
	// if g.state == stateDealerTurn {
	// 	g.state++
	// 	return nil
	// }
	// if g.state == statePlayerTurn {
	// 	g.handIdx++
	// 	if g.handIdx >= len(g.player) {
	// 		g.state++
	// 	}
	// 	return nil
	// }
	// return errors.New("invalid state")
}

func endRound(g *Game, ai AI) {
	dScore := Score(g.dealer...)
	dBlackjack := Blackjack(g.dealer...)
	allHands := make([][]deck.Card, len(g.player))
	for hi, hand := range g.player {
		cards := hand.cards
		allHands[hi] = cards
		winnings := hand.bet
		pScore, pBlackjack := Score(hand.cards...), Blackjack(cards...)
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
	}
	ai.Results(allHands, g.dealer)
        // fmt.Printf("Balance: %d\n", g.balance)
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
