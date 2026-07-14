// File name: ...\s04\09_slices_string\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {
	s := "ABCDEFGHIJKL"

	slc := s[1:5]

	fmt.Printf("s=%v type(s)=%T slc=%v type(slc)=%T\n", s, s, slc, slc)
}
