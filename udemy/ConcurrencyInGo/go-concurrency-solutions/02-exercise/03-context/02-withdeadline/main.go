package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	result string
}

func main() {

	// TODO: set deadline for goroutine to return computational result.
	deadLine := time.Now().Add(52 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadLine)
	defer cancel()

	compute := func(ctx context.Context) <-chan data {
		ch := make(chan data)
		go func() {
			defer close(ch)
			deadLine, ok := ctx.Deadline()
			if ok {
				if deadLine.Sub(time.Now().Add(50*time.Millisecond)) < 0 {
					fmt.Println("not sufficient time given... terminating")
					return
				}
			}
			// Simulate work.
			time.Sleep(50 * time.Millisecond)

			// Report result.
			select {
			case ch <- data{"123"}:
			case <-ctx.Done():
				return
			}
		}()
		return ch
	}

	// Wait for the work to finish. If it takes too long move on.
	ch := compute(ctx)
	d, ok := <-ch
	if ok {
		fmt.Printf("work complete: %s\n", d)
	}

}
