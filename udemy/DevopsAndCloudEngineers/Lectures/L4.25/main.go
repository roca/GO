package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("one")
	c := make(chan string)
	go testFunction(c)
	fmt.Println("two")
	time.Sleep(2 * time.Second)
	areWeFinished := <-c
	fmt.Println("Finished:", areWeFinished)
}

func testFunction(c chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println("checking...")
		time.Sleep(1 * time.Second)
	}
	c <- "yes we are !"
}
