package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	args := os.Args

	go boring(args[1])
	fmt.Println("I'm listening.")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring: I'm leaving.")
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}
