// File name: ...\s12\07_channel_multipe_receivers\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"time"
)

var wordSet = []string{"red", "green", "blue", "yellow", "white"}

func main() {

	receivers := []interface{}{receiver1, receiver2}

	var c = make(chan string)
	var b = make(chan bool)

	go sender(c)

	for i := range receivers {
		go receivers[i].(func(chan bool, chan string))(b, c)
	}

	for range receivers {
		<-b
	}
}

func sender(c chan string) {
	for _, w := range wordSet {
		c <- w
	}
	close(c)
}

func receiver1(b chan bool, c chan string) {
	for w := range c {
		fmt.Printf("(r1) %v \n", w)
		time.Sleep(2 * time.Second)
	}

	b <- true
}

func receiver2(b chan bool, c chan string) {
	for w := range c {
		fmt.Printf("(r2) %v \n", w)
		time.Sleep(time.Second)
	}
	b <- true
}
