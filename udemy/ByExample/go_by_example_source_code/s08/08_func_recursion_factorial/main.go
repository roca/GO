// File name: ...\s08\08_func_recursion_factorial\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	fmt.Println("=>", factorial(3))
	fmt.Println("=>", factorial(4))
	fmt.Println("=>", factorial(7))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	clear(64)
	fmt.Print(&n, " ") //for debugging purposes
	return n * factorial(n-1)
}

func clear(i int) []byte {
	if i == 0 {
		return nil
	}
	return clear(i - 1)
}
