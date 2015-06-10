// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/ExxRWe6jUz

// Sample program to show how constants do have a parallel type system.
package main

import "fmt"

const (
	// Max integer value on 64 bit architecture.
	maxInt = 9223372036854775807

	// Much larger value than int64.
	bigger = 9223372036854775808543522345

	// Will NOT compile
	// biggerInt int64 = 9223372036854775808543522345
)

// main is the entry point for the application.
func main() {
	fmt.Println("Will Compile")
}
