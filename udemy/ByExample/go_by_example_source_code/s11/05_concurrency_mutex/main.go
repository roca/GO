// File name: ...\s11\05_concurrency_mutex\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var waitG sync.WaitGroup
var counter = 0
var m sync.Mutex

func main() {
	waitG.Add(2)
	go numbers(1)
	go numbers(2)
	waitG.Wait()

	fmt.Printf("\ncounter: %d", counter)
}

func numbers(callID int) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 10; i++ {

		time.Sleep(200 * time.Millisecond)

		m.Lock()
		counter++
		fmt.Printf("(%d) %d %d\n", callID, rand.Intn(20)+20, counter)
		m.Unlock()
	}
	waitG.Done()
}
