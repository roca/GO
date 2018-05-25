package raindrops

import (
	"fmt"
	"strings"
)

const testVersion = 2

// Convert tranforms int to custom phrase
// in keeping with rules in README
func Convert(n int) string {
	factors := [3]int{3, 5, 7}
	labels := [3]string{"Pling", "Plang", "Plong"}

	phrase := fmt.Sprintf("%d", n)
	for i, factor := range factors {
		if n%factor == 0 {
			phrase += labels[i]
		}
	}

	aPhrase := strings.Split(phrase, "P")
	if len(aPhrase) == 1 {
		phrase = aPhrase[0]
	} else {
		phrase = fmt.Sprintf("P%s", strings.Join(aPhrase[1:], "P"))
	}
	return phrase
}
