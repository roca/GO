// File name: ...\s02\05_binary_decimal_etc\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d %o %x %b \n", byte('A'), byte('A'), byte('A'), byte('A'))

	fmt.Printf("%d %[1]o %[1]x %[1]b \n", byte('A'))

	fmt.Printf("%d %o %x %b \n", 'A', 'A', 'A', 'A')
}
