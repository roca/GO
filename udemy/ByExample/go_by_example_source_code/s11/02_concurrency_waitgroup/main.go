// File name: ...\s11\02_concurrency_waitgroup\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var waitG sync.WaitGroup

func main() {
	waitG.Add(2)
	go numbers()
	go alphabets()
	waitG.Wait()

	// time.Sleep(3200 * time.Millisecond)
	fmt.Println("\nmain terminated")
}

func numbers() {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 15; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d ", rand.Intn(20)+20) // formula: rand.Intn(max-min) + min
	}
	waitG.Done()
}

func alphabets() {
	for i := 'C'; i <= 'G'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
	waitG.Done()
}
