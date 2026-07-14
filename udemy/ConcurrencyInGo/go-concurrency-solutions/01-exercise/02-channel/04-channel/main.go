package main

import "fmt"

// TODO: Implement relaying of message with Channel Direction

func genMsg(message string, ch1 chan<- string) {
	// send message on ch1
	ch1 <- message
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	// recv message on ch1
	message := <-ch1
	// send it on ch2
	ch2 <- message
}

func main() {

	var biDirectional chan string
	biDirectional = make(chan string)

	var ch1 chan<- string
	var ch2 <-chan  string

	// create ch1 and ch2
	ch1 = biDirectional
	ch2 = biDirectional

	// spine goroutine genMsg and relayMsg
	go genMsg("message", ch1)

	go relayMsg(ch2, ch1)

	// recv message on ch2
	fmt.Println(<-ch2)
}
