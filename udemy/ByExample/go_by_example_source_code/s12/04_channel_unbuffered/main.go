// File name: ...\s12\04_channel_unbuffered\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"sync"
	"time"
)

var waitG sync.WaitGroup

func main() {
	c := make(chan int)

	waitG.Add(2)
	go send(c)
	go receive(c)
	waitG.Wait()

	// time.Sleep(2 * time.Second)
	time.Sleep(3 * time.Millisecond)

}

func send(c chan int) {
	for i := 1; i < 6; i++ {
		c <- i
	}
	waitG.Done()
}

func receive(c chan int) {
	for i := 1; i < 6; i++ {
		// for {
		fmt.Println(<-c)
	}
	// with the infinite loop, the No waitG.Done() won't be reached.
	waitG.Done()
}
