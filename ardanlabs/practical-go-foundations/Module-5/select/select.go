package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan string, 1), make(chan string, 1)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "one"
	}()
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "two"
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	select {
	case v := <-ch1:
		fmt.Println("ch1:", v)
	case v := <-ch2:
		fmt.Println("ch2:", v)
	// case <-time.After(10 * time.Millisecond):
	// 	fmt.Println("timeout")
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}
