// File name: ...\s02\23_utf8_5\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {

	s := "こんにちは"

	fmt.Printf("% x\n\n", s)

	r := []rune(s)
	fmt.Printf("%x\n", r)

	fmt.Println(string(r))
}
