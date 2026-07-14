// File name: ...\s04\15_slices_change_size\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	var x []int
	xLen := len(x)

	fmt.Println(x, len(x), cap(x))

	a := []int{100, 200, 300}

	if xLen <= len(a) {
		xLen := len(a)
		xCap := xLen + 1

		x = make([]int, xLen, xCap)
		copy(x, a)
	}

	fmt.Println(a, len(a), cap(a))
	fmt.Println(x, len(x), cap(x))
}
