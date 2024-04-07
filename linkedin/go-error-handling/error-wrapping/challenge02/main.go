/// Write your answer here, and then test your code.

package main

import (
	"fmt"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

type EvenNumberError struct {
	Number int
}

func (e *EvenNumberError) Error() string {
	return fmt.Sprintf("%d is not an odd number", e.Number)
}

func isOdd(n int) error {
	// Your code goes here.
	if n%2 == 0 {
		return &EvenNumberError{Number: n}
	}
	return nil
}

func main() {
	var result string
	for i := 0; i < 5; i++ {
		err := isOdd(i)
		if err != nil {
			result += fmt.Sprintf("Error: %v\n", err)
		} else {
			result += fmt.Sprint("Number is odd\n")
		}
	}
	fmt.Println(result)
}
