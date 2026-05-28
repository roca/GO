package main

const (
	DeckSize = 52
)

type CardCounter struct {
	SeenCards      map[string]int // Map to keep track of cards by value
	RunningCount   int            // Running count for hi-lo strategy
	TrueCount      float64        // True count (running count / cards remaining)
	DecksRemaining float64        // Estimate of remaining cards
}

func NewCardCounter() *CardCounter {
	counter := &CardCounter{
		SeenCards:      make(map[string]int),
		RunningCount:   0,
		TrueCount:      0.0,
		DecksRemaining: 1.0,
	}

	// Initialize counts for each card value to 0
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	for _, value := range values {
		counter.SeenCards[value] = 0
	}

	return counter
}

func (cc *CardCounter) Reset() {
	cc.RunningCount = 0
	cc.TrueCount = 0.0
	cc.DecksRemaining = 1.0

	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	for _, value := range values {
		cc.SeenCards[value] = 0
	}
}

// TrackCard updates the card counter's state based on a newly seen card
func (cc *CardCounter) TrackCard(card Card) {
	// Update seen cards count
	cc.SeenCards[card.Value]++

	// Update running count based on hi-lo strategy
	switch card.Value {
	case "2", "3", "4", "5", "6":
		cc.RunningCount++
	case "10", "J", "Q", "K", "A":
		cc.RunningCount--
	}

	// Update decks remaining estimation (52 cards in a deck)
	totalSeen := 0
	for _, count := range cc.SeenCards {
		totalSeen += count
	}
	cc.DecksRemaining = (52.0 - float64(totalSeen)) / DeckSize
	if cc.DecksRemaining < 0.1 {
		cc.DecksRemaining = 0.1 // Avoid division by too small number
	}

	// Calculate true count
	cc.TrueCount = float64(cc.RunningCount) / cc.DecksRemaining

}

func (cc *CardCounter) ChanceOfBusting(playerScore int) float64 {
	//  If player has 21 or more, they'll bust on any hit
	if playerScore >= 21 {
		return 1.0
	}

	// Calculate how many points until the player busts
	pontsUntilBust := 21 - playerScore

	// Count unseen cards that would cause a bust
	bustCards := 0
	totalUnseenCards := 0

	// For each card value, check if it would caust a bust
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	scores := []int{11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}

	for i, value := range values {
		// Each value appears 4 times
		totalInDeck := 4
		seen := cc.SeenCards[value]
		unseen := totalInDeck - seen
		unseen = max(0, unseen) // Ensure we don't count negative unseen cards

		totalUnseenCards += unseen

		// If this  card would cause a bust, count it
		if scores[i] > pontsUntilBust {
			bustCards += unseen
		}
	}

	// Avoid division by zero
	if totalUnseenCards == 0 {
		return 0.5 // Default to 50% if we somehowq have no unseen cards
	}

	return float64(bustCards) / float64(totalUnseenCards)
}

func (cc *CardCounter) DealerChanceOfBusting(dealerUpCard Card) float64 {

	// Base probabilities of dealer busting based on their up card (simplified)
	bustProbabilities := map[string]float64{
		"A":  0.17,
		"2":  0.35,
		"3":  0.37,
		"4":  0.40,
		"5":  0.42,
		"6":  0.42,
		"7":  0.26,
		"8":  0.24,
		"9":  0.23,
		"10": 0.21,
		"J":  0.21,
		"Q":  0.21,
		"K":  0.21,
	}

	// Adjust base probability based on our card counting information
	baseProbability := bustProbabilities[dealerUpCard.Value]

	// If the true count is positive (more low cards have been seen),
	// the dealer is more likely to bust.
	adjustment := cc.TrueCount * 0.02 // Adjust by 2% per true count point

	return baseProbability + adjustment
}
