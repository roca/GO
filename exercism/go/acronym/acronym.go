package acronym

import (
	"regexp"
	"strings"

	"github.com/fatih/camelcase"
)

const testVersion = 1

func abbreviate(phrase string) string {
	rgx := regexp.MustCompile(" +|-|:|,")
	words := rgx.Split(phrase, -1)
	abbreviation := ""
	for _, word := range words {
		camelCasedWords := camelcase.Split(word)
		if len(camelCasedWords) == 1 {
			abbreviation += strings.Split(strings.ToUpper(word), "")[0]
		} else {
			for _, word := range camelCasedWords {
				abbreviation += strings.Split(strings.ToUpper(word), "")[0]
			}
		}
	}
	return abbreviation

}
