// File name: ...\s12\06_channel_semaphore\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

var wordSet1 = []string{"small", "medium", "large"}
var wordSet2 = []string{"red", "green", "blue", "yellow"}

var c = make(chan string)
var b = make(chan bool)

func main() {

	f := []interface{}{sender1, sender2}

	for i := range f {
		go f[i].(func())()
	}
	go closeSenders()

	fmt.Println("Before getting to the 'channel for range loop'.")
	for val := range c {
		fmt.Println(val)
	}
}

func sender1() {
	for _, w := range wordSet1 {
		c <- w
	}
	b <- true
}

func sender2() {
	for _, w := range wordSet2 {
		c <- w
	}
	b <- true
}

func closeSenders() {
	<-b
	<-b
	close(c)
	// close(b)		//better to close the channel 'b' as well!
}
