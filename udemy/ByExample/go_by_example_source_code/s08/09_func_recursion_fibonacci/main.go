// File name: ...\s08\09_func_recursion_fibonacci\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

// ASSIGNMENT: In mathematics, the Fibonacci numbers are the numbers in the following
// integer sequence, called the Fibonacci sequence, and characterized by the fact
// that every number after the first two is the sum of the two preceding ones:
// 0, 1 , 1 , 2 , 3 , 5 , 8 , 13 , 21 , 34 , 55 , 89 , 144 , ...
// Implement the Fibonacci series using recursion.
// Ref: https://en.wikipedia.org/wiki/Fibonacci_number

func main() {
	var i int
	for i = 0; i < 15; i++ {
		fmt.Printf("%d ", fibonacci(i)) //0 1 1 2 3 5 8 13 21 34 55 89 144 233 377
	}

	fmt.Printf("\n\n%d ", fibonacci(4))
}

func fibonacci(n int) int {

	//fmt.Print(n, " ") //for debugging purposes
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
