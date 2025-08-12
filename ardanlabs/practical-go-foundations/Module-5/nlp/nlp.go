package nlp

import (
	"regexp"
	"strings"

	"github.com/roca/GO/tree/staging/ardanlabs/practical-go-foundations/Module-5/nlp/stemmer"
)

var (
	// "Who's on first?" -> [Who s on first]
	wordRe = regexp.MustCompile(`[a-zA-Z]+`)
)

// Tokenize returns tokens (lower case) found in text.
func Tokenize(text string) []string {

	words := wordRe.FindAllString(text, -1)
	var tokens []string
	for _, w := range words {
		token := strings.ToLower(w)
		token = stemmer.Stem(token)
		if token != "" {
			tokens = append(tokens, token)
		}
	}
	return tokens
}
