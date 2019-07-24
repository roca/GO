/*
Package etl implements Transform which transforms
a scrabble list of letters per score to a score per letter
*/
package etl

import "strings"

var charScores = map[string]int{
	"a": 1, "e": 1, "i": 1, "o": 1, "u": 1, "l": 1, "n": 1, "r": 1, "s": 1, "t": 1,
	"d": 2, "g": 2,
	"b": 3, "c": 3, "m": 3, "p": 3,
	"f": 4, "h": 4, "v": 4, "w": 4, "y": 4,
	"k": 5,
	"j": 8, "x": 8,
	"q": 10, "z": 10,
}

/*
Transform which transforms a scrabble list of letters per score
to a score per letter
*/
func Transform(input map[int][]string) map[string]int {
	output := map[string]int{}

	for _, chars := range input {
		for _, char := range chars {
			c := strings.ToLower(char)
			output[c] = charScores[c]
		}
	}
	return output
}
