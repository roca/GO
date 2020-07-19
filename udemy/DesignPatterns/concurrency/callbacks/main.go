package main

import (
	"fmt"
	"strings"
	"sync"
)

var wait sync.WaitGroup

func toUpperSync(word string, f func(string)) {
	go f(strings.ToUpper(word))
}

func main() {

	wait.Add(2)
	toUpperSync("Hello Callbacks!", func(v string) {
		toUpperSync("Hello Callbacks!", func(v string) {
			fmt.Printf("Callbacks: %s\n", v)
			wait.Done()
		})
	})
	println("Waiting async response...")
	wait.Wait()
}
