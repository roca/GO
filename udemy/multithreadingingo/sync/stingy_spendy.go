package main

import (
	"sync"
)

var (
	money = 100
	wg    sync.WaitGroup
	//lock  = sync.Mutex{}
)

func stingy() {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		//lock.Lock()
		money += 10
		//lock.Unlock()

	}
	println("Stingy Done")
}

func spendy() {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		//lock.Lock()
		money -= 10
		//lock.Unlock()

	}
	println("Spendy Done")
}

func main() {

	for i := 1; i <= 400; i++ {
		wg.Add(2)
		go stingy()
		go spendy()
		wg.Wait()
		print(money, "\n")
	}

}
