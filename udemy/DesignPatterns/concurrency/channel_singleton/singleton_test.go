package channel_singleton

import (
	"fmt"
	"testing"
	"time"
)

func TestStartInstance(t *testing.T) {
	singleton := GetInstance()
	singleton2 := GetInstance()

	n := 100000

	for i := 0; i < n; i++ {
		go singleton.AddOne()
		go singleton2.AddOne()
	}

	fmt.Printf("Before loop, current count is %d\n", singleton.GetCount())

	var val int
	for val != n*2 {
		val = singleton.GetCount()
		time.Sleep(10 * time.Millisecond)
	}

	singleton.Stop()
}
