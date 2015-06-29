// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Create a package named toy with a single exported struct type named Toy. Add
// the exported fields Name and Weight. Then add two unexported fields named
// onHand and sold. Declare a factory function called New to create values of
// type toy and accept parameters for the exported fields. Then declare methods
// that return and update values for the unexported fields.
//
// Create a program that imports the toy package. Use the New function to create a
// value of type toy. Then use the methods to set the counts and display the
// field values of that toy value.
package main

import (
	"fmt"

	"github.com/ArdanStudios/gotraining/04-packaging_exporting/exercises/exercise1/toy"
)

// main is the entry point for the application.
func main() {
	// Create a value of type toy.
	toy := toy.New("Bat", 28)

	// Update the counts.
	toy.UpdateOnHand(100)
	toy.UpdateSold(2)

	// Display each field separately.
	fmt.Println("Name", toy.Name)
	fmt.Println("Weight", toy.Weight)
	fmt.Println("OnHand", toy.OnHand())
	fmt.Println("Sold", toy.Sold())
}
