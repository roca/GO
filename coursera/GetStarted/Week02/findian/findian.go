package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var x string

	fmt.Println("Please input enter a string")
	fmt.Scan(&x)

	// Trim away Leading and Trailing spaces
	x = strings.TrimSpace(x)

	switch Found(x, "(?i)^i.*a.*n$") {
	case true:
		fmt.Println("Found!")
	default:
		fmt.Println("Not Found!")
	}
}

// Found ..
func Found(s, pattern string) bool {
	regex, _ := regexp.Compile(pattern)
	return regex.MatchString(s)
}
