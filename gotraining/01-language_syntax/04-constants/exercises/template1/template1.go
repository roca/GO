// Declare an untyped and typed constant and display their values.
//
// Multiply two literal constants into a typed variable and display the value.
package main

import "fmt"

// Add imports.

// Declare a constant of kind string and assign a value.
const s1 string = "hello"

// Declare a constant of type integer and assign a value.
const i int = 100
const f float64 = 3.9

// main is the entry point for the application.
func main() {
	// Display the value of both constants.
	fmt.Println(s1, i)
	// Divide a constant of kind integer and kind floating point and
	// assign the result to a variable.
	// Display the value of the variable.
	x := 1 / f
	fmt.Println(x)
}
