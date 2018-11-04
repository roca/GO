// File name: ...\s11\01_concurrency_goroutine\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// First run without 'go' (no concurrency); then with 'go' (concurrency);
	go numbers()
	go alphabets()

	time.Sleep(3200 * time.Millisecond)
	fmt.Println("\nmain terminated")
}

func numbers() {
	// func Now() Time
	//   returns the current local time; requires "time" package
	// func (t Time) UnixNano() int64
	//   returns 't' as a Unix time, the number of nanoseconds elapsed
	//   since January 1, 1970 UTC; requires "time")
	//
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 15; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d ", rand.Intn(20)+20) // formula: rand.Intn(max-min) + min
	}
}

func alphabets() {
	for i := 'C'; i <= 'G'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
