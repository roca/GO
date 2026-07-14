package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fanIn(inputs ...<-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-inputs[0]
		}
	}()

	go func() {
		for {
			c <- <-inputs[1]
		}
	}()

	return c
}

func main() {

	c := fanIn(boring("Joe"), boring("Ann")) // Function returning a channel

	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("You're both boring: I'm leaving.")

}

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() { // We launch the goroutine from inside the function
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}
