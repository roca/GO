package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println(n)
			wg.Done()
		}(i)
	}

	wg.Wait()

	// time.Sleep(10 * time.Millisecond)

	ch := make(chan string)

	go func() { ch <- "hi" }() // send
	msg := <-ch                // receive
	fmt.Println(msg)

	go func() {
		for i := 0; i < 10; i++ {
			msg := fmt.Sprintf("message #%d", i+1)
			ch <- msg
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}

	/* for/range loop with ok
		for {
			msg, ok := <-ch
			if !ok {
				break
			}
	}		fmt.Println(msg)
	        }
	*/

	msg = <-ch // ch is closed, so this will return zero value of string
	fmt.Printf("closed: zero value for %T is %#v\n", msg, msg)

	msg, ok := <-ch // ch is closed, so this will return zero value of string
	fmt.Printf("closed: zero value for %T is %#v (ok=%#v)\n", msg, msg, ok)

	// ch <- "hello" // panic: send on closed channel
	// close(ch) // panic: close of closed channel

	// send/receive to a nil channel will block forever

	fmt.Println("sleepSort:", sleepSort([]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}))
}

/*

For every value "n" in values, spin a goroutine that will
- sleep "n" milliseconds
- Send "n" over a channel

In the function body, collect values from the channel to a slice and return it.
*/

func sleepSort(values []int) []int {
	// wg := sync.WaitGroup{}
	ch := make(chan int)
	var sorted []int

	for n := range values {
		// wg.Add(1)
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n
			// wg.Done()
		}(n)
	}
	// wg.Wait()

	for n := range ch {
		sorted = append(sorted, n)
	}

	return sorted
}
