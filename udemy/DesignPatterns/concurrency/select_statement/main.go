package main

import "time"

func main() {
	helloCh := make(chan string, 1)
	goodbyeCh := make(chan string, 1)
	quitCh := make(chan bool)
	go receiver(helloCh, goodbyeCh, quitCh)
	go sendString(helloCh, "hello!")
	time.Sleep(time.Second)
	go sendString(goodbyeCh, "goodbye!")
	<-quitCh
}

func sendString(ch chan<- string, s string) {
	ch <- s
}

func receiver(helloCh, goodbyeCh <-chan string, quitCh chan<- bool) {
	for {
		select {
		case msg := <-helloCh:
			println(msg)
		case msg := <-goodbyeCh:
			println(msg)
		case <-time.After(time.Second * 2):
			println("Nothing received in 2 seconds. Exiting")
			quitCh <- true
			break
		}
	}
}
