package main

import (
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	value int
}

func main() {
	counter := Counter{}
	for i := 0; i < 10; i++ {
		go func(i int) {
			counter.Lock()
			counter.value++
			defer counter.Unlock()
		}(i)
	}
	time.Sleep(time.Second)
	counter.Lock()
	println(counter.value)
	defer counter.Unlock()
}
