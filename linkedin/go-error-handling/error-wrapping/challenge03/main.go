// Write your answer here, and then test your code.
// Your job is to implement the isOdd() method.

package main

import (
	"errors"
	"fmt"

	"log"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

func isOdd(n int) error {
	// Your code goes here.
	if n%2 == 0 {
		err := errors.New("Logged error: even number")
		log.Println(err)
		return err
	}
	return nil
}

func main() {
	var result string
	for i := 0; i < 5; i++ {
		err := isOdd(i)
		if err != nil {
			result += fmt.Sprintf("%v\n", err)
		} else {
			result += fmt.Sprintf("%d is odd.\n", i)
		}
	}
	fmt.Println(result)
}
