/// Write your answer here, and then test your code.

package main

import "fmt"

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

func isOdd(n int) error {
	if n%2 == 0 {
		return fmt.Errorf("Number %d is even", n)
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
			result += fmt.Sprintf("Number is odd\n")
		}
		fmt.Println(result)
	}
}
