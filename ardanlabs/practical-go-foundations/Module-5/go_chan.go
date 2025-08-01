package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := range 3 {
		go func() {
			fmt.Println("goroutine:", i)
		}()
	}

	time.Sleep(10 * time.Millisecond)

	ch := make(chan int)
	var v int

	go func() {
		ch <- 7 //send
	}()

	v = <-ch
	//v := <-ch // receive
	fmt.Println(v)
}
