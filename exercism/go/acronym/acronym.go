package acronym

import (
	"regexp"
	"strings"

	"github.com/fatih/camelcase"
)

const testVersion = 1

func abbreviate(input string) string {
	a := regexp.MustCompile(" +|-|:|,")
	inputs := a.Split(input, -1)
	abbreviation := ""
	for _, word := range inputs {
		camelCaseWords := camelcase.Split(word)

		if len(camelCaseWords) == 1 {
			abbreviation += strings.Split(strings.ToUpper(word), "")[0]
		} else {
			for _, word := range camelCaseWords {
				abbreviation += strings.Split(strings.ToUpper(word), "")[0]
			}
		}
	}
	return abbreviation

}
