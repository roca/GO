// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/dJk2eycWhH

// Sample program to show how to use a third index slice.
package main

import (
	"fmt"
)

// main is the entry point for the application.
func main() {
	// Create a slice of strings with different types of fruit.
	slice := []string{"Apple", "Orange", "Banana", "Grape", "Plum"}
	inspectSlice(slice)

	// Take a slice of slice. We want just index 2
	// takeOne[0] = "Banana"
	// Length:   3 - 2
	// Capacity: 5 - 2
	takeOne := slice[2:3]
	inspectSlice(takeOne)

	// For slice[ i : j : k ] the
	// Length:   j - i
	// Capacity: k - i

	// Take a slice of just index 2 with a length and capacity of 1
	// takeOneCapOne[0] = "Banana"
	// Length:   3 - 2
	// Capacity: 3 - 2
	takeOneCapOne := slice[2:3:3] // Use the third index position to
	inspectSlice(takeOneCapOne)   // set the capacity to 1.

	// Append a new element which will create a new
	// underlying array to increase capacity.
	takeOneCapOne = append(takeOneCapOne, "Kiwi")
	inspectSlice(takeOneCapOne)
}

// inspectSlice exposes the slice header for review.
func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for index, value := range slice {
		fmt.Printf("[%d] %p %s\n",
			index,
			&slice[index],
			value)
	}
}
