// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/I-gUNt3biw

// Basic command line program that accepts arguments to websites and retrieves
// them and record response times.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Parse all arguments
	flag.Parse()

	// flag.Args contains all non-flag arguments
	sites := flag.Args()

	// Lets keep a reference to when we started
	start := time.Now()

	for _, site := range sites {
		// start a timer for this request
		begin := time.Now()

		// Retreive the site
		if _, err := http.Get(site); err != nil {
			fmt.Println(site, err)
			continue
		}

		fmt.Printf("Site %q took %s to retrieve.\n", site, time.Since(begin))
	}

	fmt.Printf("Entire process took %s\n", time.Since(start))
}
