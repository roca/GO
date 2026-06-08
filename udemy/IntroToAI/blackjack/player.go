package main

type Player struct {
	Name   string
	Hand   []Card
	Score  int
	IsAI   bool
	IsBust bool
}

func NewPlayer(name string, isAI bool) Player {
	return Player{
		Name:   name,
		Hand:   []Card{},
		Score:  0,
		IsAI:   isAI,
		IsBust: false,
	}
}

// CalulateScore calculates the player's score based on their hand of cards.
func (p *Player) CalculateScore() int {
	// First, add up all the non-ace cards.
	nonAceScore := 0
	aces := 0

	// Go through each card
	for _, card := range p.Hand {
		if card.Value == "A" {
			// Count the number of aces
			aces++
		} else {
			// Add the value of non-ace cards
			nonAceScore += card.Score
		}
	}

	// Handle aces if any
	acrScore := 0
	for range aces {
		// If adding 11 keeps the score <= 21, use 11, otherwise use 1
		if nonAceScore+acrScore+11 <= 21 {
			acrScore += 11
		} else {
			acrScore += 1
		}
	}

	return nonAceScore + acrScore
}

// AddCard adds a card to the player's hand and updates their score.
func (p *Player) AddCard(card Card, cardCounter *CardCounter) {
	// Add the card to the player's hand
	p.Hand = append(p.Hand, card)

	// Update the player's score
	p.Score = p.CalculateScore()

	// If we are keeping track of cards, update the card cardCounter
	if cardCounter != nil {
		cardCounter.TrackCard(card)
	}
}
