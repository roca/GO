// Package scrabble implements Score which returns a scrabble number score for a given word
package scrabble

import (
	"strings"
	"sync"
)

// Score returns the scrabble number score for a given word
func Score(input string) int {
	var wg sync.WaitGroup
	wg.Add(len(input))

	score := 0
	word := strings.ToUpper(input)

	for _, c := range word {
		go func(char byte) {
			defer wg.Done()
			score += CharScore(char)
		}(byte(c))
	}

	wg.Wait()
	return score
}

// CharScore returns the scrabble number score for a given character
func CharScore(char byte) int {
	charScores := map[string]int{
		"AEIOULNRST": 1,
		"DG":         2,
		"BCMP":       3,
		"FHVWY":      4,
		"K":          5,
		"JX":         8,
		"QZ":         10,
	}

	for key, value := range charScores {
		if strings.IndexByte(key, char) != -1 {
			return value
		}
	}

	return 0
}
