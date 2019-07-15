// Package proverb implements Proverb which returns a proverb given a list of inputs as []string
package proverb

import (
	"fmt"
)

// Proverb returns a proverb given a list of inputs as []string
func Proverb(rhyme []string) []string {

	var proverb []string

	for i := range rhyme {
		if i != 0 {
			proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]))
		}
	}
	if len(rhyme) != 0 {
		proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	}

	return proverb
}
