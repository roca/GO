// Clock stub file

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

import (
	"fmt"
	"time"
)

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.

type Clock struct {
	currentTime time.Time
} // Complete the type definition.  Pick a suitable data type.

func New(hour, minute int) Clock {

	midNight := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)

	c := Clock{currentTime: midNight}
	c.currentTime = c.currentTime.Add(time.Duration(hour) * time.Hour)
	c.currentTime = c.currentTime.Add(time.Duration(minute) * time.Minute)
	return c
}

func (c Clock) String() string {

	return fmt.Sprintf("%02d:%02d", c.currentTime.Hour(), c.currentTime.Minute())
}

func (c Clock) Add(minutes int) Clock {
	c.currentTime = c.currentTime.Add(time.Duration(minutes) * time.Minute)
	return c
}

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.
