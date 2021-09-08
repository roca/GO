package module01

import (
	"fmt"
	"strings"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	s := []string{}
	for i := 1; i <= n; i++ {
		s = append(s, FizzBuzzValue(i))
	}
	fmt.Println(strings.Join(s, ", "))
}

func FizzBuzzValue(i int) string {
	switch {
	case i%3 == 0 && i%5 == 0:
		return "Fizz Buzz"
	case i%3 == 0:
		return "Fizz"
	case i%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", i)
	}
}
