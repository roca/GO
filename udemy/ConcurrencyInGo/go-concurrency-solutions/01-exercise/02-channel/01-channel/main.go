package main

import "fmt"

func main() {
	ch := make(chan int)
	go func(a, b int) {
		ch <- a + b
	}(1, 2)
	// get the value computed from goroutine
	fmt.Printf("computed value %v\n", <-ch)
}
