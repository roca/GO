package main

import (
	"fmt"
	"time"
)

func AdvancedAIDecision(player Player, dealerUpCard Card, cardCounter *CardCounter) string {
	score := player.Score

	// Always going to stand on 19 or higher
	if score >= 19 {
		return Stand
	}

	// Calculate bust probability if we hit
	bustProbability := cardCounter.ChanceOfBusting(score)

	// Calculate dealer bust probability
	dealerBustProbability := cardCounter.DealerChanceOfBusting(dealerUpCard)

	// Display AI thinking
	fmt.Printf("AI thinking: Score %d, True Count %.1f%%, Bust Probability %.1f%%, Dealer Bust Probability %.1f%%\n",
		score,
		cardCounter.TrueCount,
		bustProbability*100,
		dealerBustProbability*100,
	)
	time.Sleep(500 * time.Millisecond) // Simulate thinking time

	// Decision logic incorporating card counting
	if score >= 17 {
		// With 17 or 18, consider the true count
		if cardCounter.TrueCount > 0 {
			// Positive means more high cards left, so stand
			return Stand
		} else if dealerUpCard.Score >= 7 && cardCounter.TrueCount < -2 {
			// Against a strong dealer with a very negative count, so hit on 17
			if score == 17 {
				return Hit
			}
			return Stand
		}
		return Stand
	}

	// Soft hands (have an ace counted as 11)
	hasAce := false
	for _, card := range player.Hand {
		if card.Value == "A" && score <= 21 {
			hasAce = true
			break
		}
	}

	if hasAce {
		// Soft 18 or higher
		if score >= 18 {
			// Stand unless dealer has a 9, 10, or Ace with a negative count
			if (dealerUpCard.Score >= 9 || dealerUpCard.Value == "A") && cardCounter.TrueCount < -1 {
				return Hit
			}
			return Stand
		}
		// Soft 17 or lower, hit
		return Hit
	}

	// 16 or loweer with considerations
	if score <= 16 {
		// If low risk of busting and dealer is likely to bust, then stand
		if bustProbability < 0.3 && dealerBustProbability > 0.4 && score >= 13 {
			return Stand
		}

		// Stand on 12-16 against dealer's 2-6 (unless the true count is very negative)
		if score >= 12 && dealerUpCard.Score >= 2 && dealerUpCard.Score <= 6 && cardCounter.TrueCount > -3 {
			return Stand
		}

		// Otherwise, hit
		return Hit
	}

	// Default to Stand if no other conditions met
	return Stand

}
