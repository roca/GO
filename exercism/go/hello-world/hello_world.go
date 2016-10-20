package hello

import "fmt"

const testVersion = 2

// HelloWorld exercise
//
func HelloWorld(input string) string {
	// Write some code here to pass the test suite.
	if input == "" {
		input = "World"
	}
	// When you have a working solution, REMOVE ALL THE STOCK COMMENTS.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, nitpickers will protest!
	return fmt.Sprintf("Hello, %v!", input)
}
