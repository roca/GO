/*
Package isogram implements IsIsogram which determines if a word or phrase is an isogram.

An isogram (also known as a "nonpattern word") 
is a word or phrase without a repeating letter, 
however spaces and hyphens are allowed to appear multiple times.
*/
package isogram

import (
	"regexp"
	"strings"
)

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(input string) bool {

	phrase := strings.ToLower(cleanUp(input))

	for _, c := range phrase {
		if strings.Count(phrase, string(c)) > 1 {
			return false
		}
	}
	return true
}

func cleanUp(input string) string {
	re := regexp.MustCompile(`\-|\s*`)
	return re.ReplaceAllLiteralString(input, "")
}
