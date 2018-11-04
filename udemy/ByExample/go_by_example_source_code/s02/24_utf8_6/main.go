// File name: ...\s02\24_utf8_6\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {

	var runes []rune

	for _, r := range "bye 再见" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	fmt.Printf("%q\n", []rune("bye 再见"))
}
