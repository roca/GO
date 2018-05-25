package gigasecond

import "time"

const testVersion = 4

func AddGigasecond(inDate time.Time) time.Time {
	return inDate.Add(time.Duration(1000000000) * time.Second)
}
