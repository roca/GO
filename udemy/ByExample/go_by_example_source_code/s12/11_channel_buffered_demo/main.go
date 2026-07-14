// File name: ...\s12\11_channel_buffered_demo\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) // test with 1 to 3

	for i := 1; i <= 3; i++ {
		go printMsg(c, i)
	}
	time.Sleep(10 * time.Second)
}

func printMsg(c chan int, id int) {
	fmt.Printf("ooo %d is waiting for a channel space...\n", id)

	c <- id
	fmt.Printf("=== %d has a channel space\n", id)

	time.Sleep(600 * time.Millisecond)
	fmt.Printf("xxx %d has released the channel space\n", id)

	<-c
}
