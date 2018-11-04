// File name: ...\s02\22_utf8_4\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello𠜎ち"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
}
