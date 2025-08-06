package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	/* Solution 1
	var mu sync.Mutex
	count := 0
	*/

	count := int64(0)

	nGR, nIter := 10, 1_000
	var wg sync.WaitGroup

	wg.Add(nGR)
	for range nGR {
		go func() {
			defer wg.Done()
			for range nIter {
				/* Solution 1
				mu.Lock()
				count++
				mu.Unlock()
				*/
				atomic.AddInt64(&count, 1)
				time.Sleep(time.Microsecond)
			}
		}()
	}

	wg.Wait()
	expected := int64(nGR*nIter) == count
	fmt.Println("Count:", count, "Expected:", expected)
}
