// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/KtrDN1BStt

// Download any document from the web and display the content in
// the terminal and write it to a file at the same time.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// main is the entry point for the application.
func main() {
	// Retrieve the RSS feed for the blog.
	resp, err := http.Get("http://www.goinggo.net/feeds/posts/default")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Arrange for the response Body to be Closed using defer.
	defer resp.Body.Close()

	// A slice of io.Writers we will write the file to.
	var writers []io.Writer

	// Send the document to stdout.
	writers = append(writers, os.Stdout)

	// Send the document to a file.
	file, err := os.Create("goinggo.rss")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Send the document to the file.
	writers = append(writers, file)

	// MultiWriter(io.Writer...) returns a single writer which multiplexes its
	// writes across all of the writers we pass in.
	dest := io.MultiWriter(writers...)

	// Write to dest the same way as before, copying from the Body.
	_, err = io.Copy(dest, resp.Body)
	if err != nil {
		fmt.Println(err)
	}
}
