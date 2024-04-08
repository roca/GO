// Write your answer here, and then test your code.
// Your job is to implement the findLargest() method.

package main

import (
	"fmt"
	"time"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// isEven() returns an error if a number is even
func isEven(n int) error {
	if n%2 == 0 {
		return fmt.Errorf("%d is an even number", n)
	}
	return nil
}

// isOdd returns a string channel containing
// whether a value is odd
func isOdd(n int) <-chan string {
	// Your code goes here.
      c := make(chan string)
      go func(n int) {
	err := isEven(n)
	if err != nil {
		c <- err.Error()
	} else {
		c <- fmt.Sprintf("%d is an odd number", n)
	}	
      }(n)

      return c
}

func main() {
	// This is how your code will be called.
	// Your answer should be the largest value in the numbers array.
	// You can edit this code to try different testing cases.
	var result string
	for i := 0; i < 5; i++ {
		result = <-isOdd(i)
		fmt.Println(result)
		time.Sleep(time.Millisecond)
	}
}
