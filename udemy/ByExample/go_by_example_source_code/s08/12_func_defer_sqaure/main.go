// File name: ...\s08\12_func_defer_sqaure\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

// ASSIGNMENT - Use the 'defer' statement to write a square function that hijacks the
// correct return type when the input parameter is 2 or 4.
// For 2 => return 6; For 4 => return 20;   formula: (x*x)+x
// For other values return (x*x)

func main() {
	fmt.Println(square2(2))
	fmt.Println(square2(4))

	fmt.Println(square2(3))
}

func square2(x int) (result int) {
	result = x * x

	defer func() {
		if x == 2 || x == 4 {
			result += x
		}
	}()

	fmt.Print("* ")
	return
}
