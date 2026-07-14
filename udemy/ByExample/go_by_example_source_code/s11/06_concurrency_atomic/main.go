// File name: ...\s11\06_concurrency_atomic\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var waitG sync.WaitGroup
var counter uint64

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

		atomic.AddUint64(&counter, 1)
		fmt.Printf("(%d) %d %d\n", callID, rand.Intn(20)+20, atomic.LoadUint64(&counter))

	}
	waitG.Done()
}
