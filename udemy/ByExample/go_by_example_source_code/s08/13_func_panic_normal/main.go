// File name: ...\s08\13_func_panic_normal\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

// This program will crash out at execution. It is to demonstrate Go's panic mechanism.
func main() {
	defer fmt.Printf("#A main\n")
	defer fmt.Printf("#B main\n")
	callPop()
}
func callPop() {
	defer fmt.Printf("#C callPop\n")
	pop(3)
}

func pop(x int) {
	fmt.Printf("pop(%d)\n", 10/x)
	defer fmt.Printf("#%d pop\n", x)
	pop(x - 1)
}
