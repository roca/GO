package main

import (
	"time"
)

func main() {
	channel := make(chan string, 1)
	go func(ch chan<- string) {
		ch <- "Hello World!"
		println("Finishing goroutine")
	}(channel)
	time.Sleep(time.Second)
	receivingCh(channel)
	// message := <-channel
	// fmt.Println(message)
}

func receivingCh(ch <-chan string) {
	msg := <-ch
	println(msg)
}
