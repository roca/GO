package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

var wg sync.WaitGroup

func scanPort(port int) {
	defer wg.Done()
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Printf("%d CLOSED (%s)\n", port, err)
		return
	}
	log.Printf("%d OPEN\n", port)
	conn.Close()
}

func main() {

	for i := 5200; i <= 5500; i++ {
		wg.Add(1)
		go func(n int) {
			scanPort(n)
		}(i)
	}

	wg.Wait()
}
