package clock

import (
	"fmt"
	"time"
)

const testVersion = 4

type Clock struct {
	currentTime time.Time
}

func New(hour, minute int) Clock {

	midNight := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)

	midNight = midNight.Add(time.Duration(hour) * time.Hour)
	midNight = midNight.Add(time.Duration(minute) * time.Minute)

	return Clock{currentTime: time.Date(0, 0, 0, midNight.Hour(), midNight.Minute(), 0, 0, time.UTC)}
}

func (c Clock) String() string {

	return fmt.Sprintf("%02d:%02d", c.currentTime.Hour(), c.currentTime.Minute())
}

func (c Clock) Add(minutes int) Clock {
	c.currentTime = c.currentTime.Add(time.Duration(minutes) * time.Minute)
	return c
}
