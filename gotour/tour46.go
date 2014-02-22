package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {

	n := 0

	return func() int {
		n++
		var a, b int = 1, 1
		for i := 2; i < n; i++ {
			a, b = b, a+b
		}
		return b
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(i, f())
	}
}
