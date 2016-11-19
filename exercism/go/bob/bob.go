package bob // package name must match the package name in bob_test.go
import "regexp"

const testVersion = 2 // same as targetTestVersion

// Responses struct for response types
type Responses struct {
	matcher *regexp.Regexp
	answer  string
}

//Hey returns a response based on input
func Hey(input string) string {

	responses := []Responses{
		{matcher: regexp.MustCompile(`\?$`), answer: "Sure."},
		{matcher: regexp.MustCompile(`[^[a-z]]*.$`), answer: "Whoa, chill out!"},
		{matcher: regexp.MustCompile(`^[a-z]*\.$`), answer: "Fine. Be that way!"},
		{matcher: regexp.MustCompile(`^\d+$`), answer: "Whatever."},
	}

	for _, response := range responses {
		if response.matcher.MatchString(input) {
			return response.answer
		}
	}

	return responses[3].answer
}
