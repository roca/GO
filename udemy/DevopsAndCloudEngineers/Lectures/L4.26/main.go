package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type myType struct {
	counter int
	mu      sync.Mutex
}

func (m *myType) increment() {
	m.mu.Lock()
	m.counter++
	m.mu.Unlock()
}

func main() {
	myTypeInstance := myType{}
	finished := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(myTypeInstance *myType) {
			myTypeInstance.mu.Lock()
			fmt.Printf("input counter: %d\n", myTypeInstance.counter)
			myTypeInstance.counter++
			//myTypeInstance.increment()
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			if myTypeInstance.counter == 5 {
				fmt.Printf("Found counter == %d\n", myTypeInstance.counter)
			}
			fmt.Printf("output counter: %d\n", myTypeInstance.counter)
			finished <- true
			myTypeInstance.mu.Unlock()
		}(&myTypeInstance)
	}
	for i := 0; i < 10; i++ {
		<-finished
	}
	fmt.Printf("Counter: %d\n", myTypeInstance.counter)
}
