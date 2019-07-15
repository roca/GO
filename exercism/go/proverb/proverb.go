// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import (
	"fmt"
)

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {

	type Nouns struct {
		First  string
		Second string
	}
	var Pairs = []Nouns{}

	var proverb []string

	for i, word := range rhyme {
		if i != 0 {
			Pairs = append(Pairs, Nouns{First: rhyme[i-1], Second: word})
		}
	}
	if len(rhyme) != 0 {
		Pairs = append(Pairs, Nouns{First: rhyme[0]})
	}

	for i, pair := range Pairs {

		if i != len(Pairs)-1 {
			proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", pair.First, pair.Second))
		} else {
			proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", pair.First))
		}

	}

	return proverb
}
