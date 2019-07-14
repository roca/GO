// Package scrabble implements Score which returns a scrabble number score for a given word
package scrabble

import "strings"

// Score returns the scrabble number score for a given word
func Score(input string) int {
	score := 0
	word := strings.ToUpper(input)
	for _, c := range word {
		score += CharScore(byte(c))
	}

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
