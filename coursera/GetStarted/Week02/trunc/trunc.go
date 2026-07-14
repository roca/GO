package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var x string

	fmt.Println("Please input a single floating point number. Include the decimal point !")
	fmt.Scan(&x)

	// Trim away Leading and Trailing spaces
	x = strings.TrimSpace(x)

	// Check that a decimal point was included
	r, _ := regexp.Compile("\\.")
	foundDecimal := r.FindString(x)
	if foundDecimal == "" {
		panic("You did not input a floating point number")
	}

	// Input should parse to a float
	f, err := strconv.ParseFloat(x, 32)
	if err != nil {
		panic("You did not input a floating point number")
	}

	fmt.Printf("The integer truncated portion of your input is %d\n", int(f))
}
