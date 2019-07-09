package phonenumber

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Number ...
func Number(input string) (string, error) {
	number, err := Format(input)
	if err != nil {
		return "", err
	}

	return number, nil
}

// Format ...
func Format(input string) (string, error) {
	re := regexp.MustCompile(`\:|\@|\!|\+|[a-z]|[A-Z]|\.|\-|\(|\)|\s*`)
	output := strings.Split(re.ReplaceAllLiteralString(input, ""), "")

	if len(output) >= 11 && output[0] != "1" {
		return "", fmt.Errorf("inccorect country code: %s", output[0])
	}

	if len(output) >= 11 {
		output = output[1:]
	}
	fmt.Printf("INPUT string: %s\t", input)
	fmt.Printf("INPUT ARRAY: %s\n", output)

	number := strings.Join(output[:], "")

	if len(number) < 10 || len(number) > 11 {
		return "", fmt.Errorf("number is inccorect length: %d", len(number))
	}

	areaCode, _ := strconv.Atoi(number[0:3])
	exchangeCode, _ := strconv.Atoi(number[3:6])

	if areaCode < 200 || exchangeCode < 200 {
		return "", fmt.Errorf("area or exchange codes < 200 %s %s", number[0:3], number[3:6])
	}

	return number, nil
}

// AreaCode ...
func AreaCode(input string) (string, error) {

	number, err := Number(input)
	if err != nil {
		return "", err
	}

	return number[0:3], nil
}
