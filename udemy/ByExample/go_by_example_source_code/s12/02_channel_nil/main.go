// File name: ...\s12\02_channel_nil\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {
	var c chan int
	// c := make(chan int) //declaring & defining using short hand declaration

	if c == nil {
		c = make(chan int)
	}

	fmt.Printf("Type of c: %T", c)
	close(c)
}
