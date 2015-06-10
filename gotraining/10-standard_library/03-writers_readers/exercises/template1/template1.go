// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/MxcJ581bt6

// Download any document from the web and display the content in
// the terminal and write it to a file at the same time.
package main

// Add imports.

// main is the entry point for the application.
func main() {
	// Download the RSS feed for "http://www.goinggo.net/feeds/posts/default".
	// Check for errors.

	// Arrange for the response Body to be Closed using defer.

	// Declare a slice of io.Writer interface values.

	// Append stdout to the slice of writers.

	// Open a file named "goinggo.rss" and check for errors.

	// Close the file when the function returns.

	// Append the file to the slice of writers.

	// Create a MultiWriter interface value from the writers
	// inside the slice of io.Writer values.

	// Write the response to both the stdout and file using the
	// MultiWriter. Check for errors.
}
