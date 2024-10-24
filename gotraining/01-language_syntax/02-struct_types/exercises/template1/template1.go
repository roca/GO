// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/ItPe2EEy9X

// Declare a struct type to maintain information about a user (name, email and age).
type user struct{
	name string
	email string
	age int

}
// Create a value of this type, initalize with values and display each field.

//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

import (
	"fmt"
)

  bill := user{
  	name: "William",
  	email: "will@gmail.com",
  	age: 55,
  }

  bill2 := struct{
	name string
	email string
	age int

}{
  	name: "William",
  	email: "will@gmail.com",
  	age: 55,
  }

// Add imports.

// Add user type and provide comment.

// main is the entry point for the application.
func main() {
	// Declare variable of type user and init using a struct literal.

	// Display the field values.

	// Declare a variable using an anonymous struct.

	// Display the field values.
}
