package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	/* Solutions #1
	var mu sync.Mutex
	count := 0
	*/

	// Solutions #2 using atomic
	var count int64

	const nG = 10
	var wg sync.WaitGroup
	for i := 0; i < nG; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10_000; j++ {
				/* Solutions #1
				mu.Lock()
				count++ // Lock the shared resource
				mu.Unlock()
				*/
				atomic.AddInt64(&count, 1) // Solutions #2
			}
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
