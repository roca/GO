// The race condition is that the output of "hello" and "world" is different for
// every run. It occurs because the go mini operating system is switching between
// the two goroutines in a non deterministic way.

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
