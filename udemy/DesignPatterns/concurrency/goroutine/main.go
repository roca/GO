package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	goRoutines := 5
	wait.Add(goRoutines)

	myFunc := func(msg string) {
		messagePrinter(msg)
		wait.Done()
	}

	for i := 0; i < goRoutines; i++ {
		go myFunc(fmt.Sprintf("ID:%d: Hello goroutines!\n", i))
	}
	wait.Wait()
}
func messagePrinter(msg string) {
	print(msg)
}
