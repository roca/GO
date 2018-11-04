// File name: ...\s04\14_slices_reference\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	var a []int

	// since slice is a reference type, it can be compared with 'nil'
	if a == nil {
		fmt.Printf("len(a)=%d\n", len(a))
	}

	b := []int{2, 3, 4}
	if b != nil {
		fmt.Printf("len(b)=%d\n", len(b))
	}

}
