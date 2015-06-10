// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/1xUWjHMB3I

// Declare three variables that are initalized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initalize the variable by
// converting the literal value of Pi (3.14).
package main

import (
	"fmt"
)

// main is the entry point for the application.
func main() {
	// Declare variables that are set to their zero value.
	var i int
	var s string

	// Display the value of those variables.
	s = "Hello"
	fmt.Printf("%s\n", s)
	fmt.Printf("%d\n", i)

	// Declare variables and initalize.
	x := "World"
	z := 7.5
	// Using the short variable declaration operator.

	// Display the value of those variables.
	fmt.Printf("%s\n", x)
	// Perform a type conversion.
	g := int32(z)
	fmt.Printf("%d\n", g)

	// Display the value of that variable.
}
