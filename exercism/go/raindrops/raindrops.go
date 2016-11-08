package raindrops

import (
	"fmt"
	"strings"
)

const testVersion = 2

// Convert tranforms int to custom message
func Convert(n int) string {
	factors := [3]int{3, 5, 7}
	phrase := fmt.Sprintf("%d", n)
	for _, factor := range factors {
		if n%factor == 0 && factor == 3 {
			phrase = phrase + "Pling"
		} else if n%factor == 0 && factor == 5 {
			phrase = phrase + "Plang"
		} else if n%factor == 0 && factor == 5 {
			phrase = phrase + "Plong"
		}
	}
	aPhrase := strings.Split(phrase, "P")
	if len(aPhrase) == 1 {
		phrase = aPhrase[0]
	} else {
		phrase = strings.Join(aPhrase[1:], "P")
	}
	return phrase
}

// The test program has a benchmark too.  How fast does your Convert convert?
