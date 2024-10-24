// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/H-h1cbBW3B

// Create a program that declares two anonymous functions. Once that counts up to
// 100 from 0 and one that counts down to 0 from 100. Display each number with an
// unique identifier for each goroutine. Then create goroutines from these functions
// and don't let main return until the goroutines complete.
//
// Run the program in parallel.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// main is the entry point for all Go programs.
func main() {
	runtime.GOMAXPROCS(2)

	// Declare a wait group and set the count to two.
	var wg sync.WaitGroup
	wg.Add(2)

	// Allocate two contexts for the scheduler to use in the
	// second part of this exercise.
	fmt.Println("Start Goroutines")

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Count down from 100 to 0.
		for count := 100; count >= 0; count-- {
			fmt.Printf("[A:%d]\n", count)
		}

		// Tell main we are done.
		wg.Done()
	}()

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Count up from 0 to 100.
		for count := 0; count < 100; count++ {
			fmt.Printf("[B:%d]\n", count)
		}

		// Tell main we are done.
		wg.Done()
	}()

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
