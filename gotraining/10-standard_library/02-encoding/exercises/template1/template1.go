// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/OkIHsVwMQ7

// Create a file with an array of JSON documents that contain a user name and email address. Declare a struct
// type that maps to the JSON document. Using the json package, read the file and create a slice of this struct
// type. Display the slice.
//
// Marshal the slice into pretty print strings and display each element.
package main

// Add imports.

// Declare a struct type named user with two fields. Name of type string and
// Email of type string. Add tags for each field for the unmarshal call.

// main is the entry point for the application.
func main() {
	// Use the os package to Open the JSON file. Check for errors.

	// Schedule the file to be closed once the function returns.

	// Declare a nil slice of user struct types.
	// Decode the JSON from the file into the slice. Check for errors.

	// Iterate over the slice and display each user value.

	// Marshal each user value and display the JSON. Check for errors.
}
