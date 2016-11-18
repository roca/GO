package bob // package name must match the package name in bob_test.go
import (
	"fmt"
	"regexp"
)

const testVersion = 2 // same as targetTestVersion

//Hey returns a response based on input
func Hey(input string) string {
	matchers := []*regexp.Regexp{
		regexp.MustCompile(`\?$`),
		regexp.MustCompile(`^[^[a-z]]*$`),
	}

	responses := []string{
		"Sure.",
		"Whoa, chill out!",
		"Fine. Be that way!",
		"Whatever.",
	}

	for idx, matcher := range matchers {
		if matcher.MatchString(input) {
			fmt.Println(idx)
			return responses[idx]
		}
	}

	return responses[3]
}
