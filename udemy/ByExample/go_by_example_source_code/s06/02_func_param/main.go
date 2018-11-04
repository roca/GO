// File name: ...\s06\02_func_param\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	fmt.Println(add(2, 21))

	x1, y1 := swap1("one", "two")
	fmt.Println(x1, y1)

	x2, y2 := swap2("one", "two")
	fmt.Println(x2, y2)
}

func add(a, b int) int {
	return a + b
}

func swap1(x, y string) (string, string) {
	return y, x
}

func swap2(x, y string) (r1, r2 string) {
	r1 = y
	r2 = x
	return
}
