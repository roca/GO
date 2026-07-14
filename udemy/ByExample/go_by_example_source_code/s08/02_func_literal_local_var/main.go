// File name: ...\s08\02_func_literal_local_var\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {

	factor := 2

	mult := func(i, j int) int {
		return i * j * factor
	}
	fmt.Println(mult(3, 4))
}
