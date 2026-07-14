package main

import (
	"sync"
	"sync/atomic"
)

var (
	money int32 = 100
	wg    sync.WaitGroup
)

func stingy() {
	for i := 1; i <= 1000; i++ {
		atomic.AddInt32(&money, 10)
		//money += 10

	}
	//println("Stingy Done")
	wg.Done()
}

func spendy() {
	for i := 1; i <= 1000; i++ {
		atomic.AddInt32(&money, -10)
		//money -= 10

	}
	//println("Spendy Done")
	wg.Done()
}

func main() {

	for i := 0; i < 10000; i++ {
		wg.Add(2)
		go stingy()
		go spendy()
		wg.Wait()
		if money != 100 {
			print("\n", i, ": ", money, "\n")
		}
	}

}
