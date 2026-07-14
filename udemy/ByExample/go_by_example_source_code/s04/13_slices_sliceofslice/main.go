// File name: ...\s04\13_slices_sliceofslice\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {
	s1 := []int{2, 3, 4}
	slices := [][]int{s1, {5, 6}, {7, 8, 9, 10}}
	var bigSlice []int

	for r := range slices {
		bigSlice = append(bigSlice, slices[r]...)
	}

	fmt.Printf("%v %T\n", slices, slices)
	fmt.Printf("%v %T\n", bigSlice, bigSlice)

	x := append(slices[0][1:], slices[2][:1]...)
	fmt.Printf("%v %T\n", x, x)
}
