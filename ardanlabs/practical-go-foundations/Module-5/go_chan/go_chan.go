package main

import (
	"fmt"
	"sync"
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

	fmt.Println(sleepSort2([]int{20, 30, 10})) // [10 20 30]

	go func() {
		for i := range 4 {
			ch <- i
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Println(">>", i)
	}

	v = <-ch
	fmt.Println("closed:", v)
	v, ok := <-ch
	fmt.Println("closed:", v, "ok", ok)
}

func sleepSort(ns []int) []int {
	var wg sync.WaitGroup
	s := []int{}

	for _, n := range ns {
		wg.Add(1)
		go func() {
			d := time.Duration(n) * time.Millisecond
			time.Sleep(d)
			s = append(s, n)
			wg.Done()
		}()
	}
	wg.Wait()
	return s
}

// collect all values from the a channel to a slice and return it
func sleepSort2(ns []int) []int {
	ch := make(chan int)

	s := []int{}

	for _, n := range ns {
		go func() {
			d := time.Duration(n) * time.Millisecond
			time.Sleep(d)
			ch <- n
		}()
	}

	i := 0
	for n := range ch {
		s = append(s, n)
		i++
		if i == len(ns) {
			close(ch)
		}
	}

	// for range ns {
	// 	s = append(s, <-ch)
	// }

	return s

}
