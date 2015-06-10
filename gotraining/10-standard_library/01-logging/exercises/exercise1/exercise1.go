// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/1-l_QmaSvy

// Setup a new program to use the log package. Set the Prefix to your first name and on each log line show
// the date and long path for the code file.
package main

import (
	"log"
	"os"
)

// init is called before main.
func init() {
	// Change the output device from the default stderr to stdout.
	log.SetOutput(os.Stdout)

	// Set the prefix string for each log line.
	log.SetPrefix("BILL: ")

	// Set the extra log info.
	setFlags()
}

// setFlags adds extra information on each log line.
func setFlags() {
	/*
	   Ldate		   // the date: 2009/01/23
	   Ltime           // the time: 01:23:23
	   Lmicroseconds   // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	   Llongfile       // full file name and line number: /a/b/c/d.go:23
	   Lshortfile      // final file name element and line number: d.go:23. overrides Llongfile
	   LstdFlags       // Ldate | Ltime // initial values for the standard logger
	*/

	log.SetFlags(log.Ldate | log.Llongfile)
}

// main is the entry point for the application.
func main() {
	log.Println("main function started")

	names := []string{"Henry", "Lisa", "Bill", "Matt"}
	log.Printf("These are named %+v\n", names)

	log.Fatalln("Terminate Program")
}
