// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/XJ4Ia1lMWl

/*
// A Duration represents the elapsed time between two instants as
// an int64 nanosecond count. The representation limits the largest
// representable duration to approximately 290 years.

type Duration int64

// Common durations. There is no definition for units of Day or larger
// to avoid confusion across daylight savings time zone transitions.

const (
        Nanosecond  Duration = 1
        Microsecond          = 1000 * Nanosecond
        Millisecond          = 1000 * Microsecond
        Second               = 1000 * Millisecond
        Minute               = 60 * Second
        Hour                 = 60 * Minute
)

// Add returns the time t+d.
func (t Time) Add(d Duration) Time
*/

// Sample program to show a idiomatic use of named types from the
// standard library and how they work in concert with other Go concepts.
package main

import (
	"fmt"
	"time"
)

// fiveSeconds is an typed constant of type Duration.
const fiveSeconds = 5 * time.Second // Duration(5) * Duration(1000000000)

// main is the entry point for the application.
func main() {
	// Use the time package to get the current date/time.
	now := time.Now()

	// Subtract 5 nanoseconds from now time using a literal constant.
	lessFiveNanoseconds := now.Add(-5)

	// Substract 5 seconds from now using a declared constant.
	lessFiveSeconds := now.Add(-fiveSeconds)

	// Display the values.
	fmt.Printf("Now     : %v\n", now)
	fmt.Printf("Nano    : %v\n", lessFiveNanoseconds)
	fmt.Printf("Seconds : %v\n", lessFiveSeconds)
}
