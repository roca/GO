package phonenumber

import (
	"fmt"
	"regexp"
	"strconv"
)

// Number ...
func Number(input string) (string, error) {
	number, err := Format(input)
	if err != nil {
		return "", err
	}

	return cleanUp(number), nil
}

// Format ...
func Format(input string) (string, error) {
	number, err := validate(input)
	if err != nil {
		return "", err
	}

	areaCode := number[0:3]
	exchangeCode := number[3:6]
	subscriberNumber := number[6:]

	intAreaCode, _ := strconv.Atoi(areaCode)
	intExchangeCode, _ := strconv.Atoi(exchangeCode)

	if intAreaCode < 200 || intExchangeCode < 200 {
		return "", fmt.Errorf("area or exchange codes was < 200:  {areacode: %s, exchangecode: %s}", number[0:3], number[3:6])
	}

	return fmt.Sprintf("(%s) %s-%s", areaCode, exchangeCode, subscriberNumber), nil
}

// AreaCode ...
func AreaCode(input string) (string, error) {

	number, err := Number(input)
	if err != nil {
		return "", err
	}

	return number[0:3], nil
}

func cleanUp(input string) string {
	re := regexp.MustCompile(`\:|\@|\!|\+|[a-z]|[A-Z]|\.|\-|\(|\)|\s*`)
	return re.ReplaceAllLiteralString(input, "")
}

func validate(input string) (string, error) {
	number := cleanUp(input)

	if len(number) >= 11 && string(number[0]) != "1" {
		return "", fmt.Errorf("inccorect country code: %s insted of 1", string(number[0]))
	}

	if len(number) >= 11 {
		number = number[1:]
	}
	// fmt.Printf("INPUT string: %s\t", input)
	// fmt.Printf("INPUT ARRAY: %s\n", output)

	if len(number) < 10 || len(number) > 11 {
		return "", fmt.Errorf("number is inccorect length: %d was not >= 10 and <= 11", len(number))
	}

	return number, nil
}
