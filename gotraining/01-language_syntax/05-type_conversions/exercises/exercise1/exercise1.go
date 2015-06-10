// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/KIdESKQc8C

// Declare a named type called counter with a base type of int. Declare a variable
// named c of type counter set to its zero value. Display the value of c.
//
// Declare a variable named c2 of type counter set to the value of 10. Display the value
// of c2.
//
// Declare a variable named i of the base type int. Attempt to assign the value
// of i to c. Does the compiler allow the assignment?
package main

import "fmt"

// counter is a named type for counting.
type counter int

// main is the entry point for the application.
func main() {
	// Declare a variable of type counter.
	var c counter

	// Display the value of c.
	fmt.Println(c)

	// Declare a second variable of type counter. Set
	// the value to 10.
	c2 := counter(10)

	// Display the value of c2.
	fmt.Println(c2)

	// Declare a variable named i of type int.
	i := 1

	// Assign the value of i to the variable named c.
	c = i

	// Does the program compile?
}
