// Write your answer here, and then test your code.
// Your job is to implement the findLargest() method.

package main

import (
	"fmt"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// isOdd() returns a panic error message if the number is not even
func isOdd(n int) {
	if n%2 != 0 {
		panic(fmt.Sprintf("%d is not even", n))
	}
}

// is even returns a boolean and prints a recover message
func isEven(n int) bool {
	b := true
	defer func() {
		if r := recover(); r != nil {
			b = false
		}
	}()
	isOdd(n)
	return b
}

func main() {
	// This is how your code will be called.
	// You can edit this code to try different testing cases.

	var answer string
	for i := 0; i < 5; i++ {
		result := isEven(i)
		answer += fmt.Sprintf("isEven() for number %d, returns %t\n", i, result)
	}

	fmt.Println(answer)
}
