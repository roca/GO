// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/d2gkKxEftw

// Declare an untyped and typed constant and display their values.
//
// Multiply two literal constants into a typed variable and display the value.
package main

import "fmt"

// server is the IP address for connecting.
const server = "124.53.24.123"

// port is the port to make that connection.
const port int16 = 9000

// main is the entry point for the application.
func main() {
	// Display the server information.
	fmt.Println(server)
	fmt.Println(port)

	// Calculate the number of minutes in 5320 seconds.
	minutes := 5320 / 60.0
	fmt.Println(minutes)
}
