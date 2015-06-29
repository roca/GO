// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/e5O_Dx5VpM

// Sample program to show how what we are doing is NOT embedding
// a type but just using a type as a field.
package main

import (
	"fmt"
)

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method that can be called via
// a value of type user.
func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n",
		u.name,
		u.email)
}

// admin represents an admin user with privileges.
type admin struct {
	person user // NOT Embedding
	level  string
}

// main is the entry point for the application.
func main() {
	// Create an admin user.
	ad := admin{
		person: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// We can acces fields methods.
	ad.person.notify()
}
