// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/lgMQHWpZul

// Declare a struct type that represents a request for a customer invoice. Include a CustomerID and InvoiceID field. Define
// tags that can be used to validate the request. Define tags that specify both the length and range for the ID to be valid.
// Declare a function named validate that accepts values of any type and processes the tags. Display the resutls of the validation.
package main

// Add imports.

// Declare a struct type named Customer. Add the fields CustomerID of type int
// with the tag `length:"3" range:"100-300"`, and field InvoiceID of type int
// with tag `length:"5" range:"60000-99999"`.

// validate performs data validation on any struct type value.
func validate( /* parameter */ ) {
	// Retrieve the value that the interface contains or points to.

	// Iterate over the fields of the struct value.
	{
		// Retrieve the field information.

		// Get the value as an int, string and the length.

		// Test the length of the value based on the tag setting.

		// Test the range of the value based on the tag setting.
	}
}

// main is the entry point for the application.
func main() {
	// Declare a variable of type Customer and initialize it.

	// Validate the value and display the results.
}
