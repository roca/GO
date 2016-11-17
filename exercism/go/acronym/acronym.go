package acronym

import (
	"regexp"
	"strings"
)

const testVersion = 1

func abbreviate(input string) string {
	a := regexp.MustCompile(" +|-")
	inputs := a.Split(input, -1)
	abbreviation := ""
	for _, word := range inputs {
		abbreviation += strings.Split(strings.ToUpper(word), "")[0]
	}
	return abbreviation

}
