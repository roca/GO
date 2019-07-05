package phonenumber

import (
	"fmt"
	"regexp"
	"strings"
)

func Number(input string) (string, error) {

	output := regexp.MustCompile(`\(|\)|\s+|-|\.|\+|\@|\!|[a-z|A-Z]+`).Split(input, -1)
	if len(output) >= 11 && output[0] != "1" {
		return "", fmt.Errorf("inccorect country code: %s", output[0])
	}

	if len(output) >= 11 {
		output = output[1:]
	}

	number := strings.Join(output[:], "")
	if len(number) < 10 || len(number) > 11 {
		return "", fmt.Errorf("number is inccorect length: %d", len(number))
	}
	return number, nil
}

func Format(input string) (string, error) {
	return "", nil
}

func AreaCode(input string) (string, error) {
	return "", nil
}
