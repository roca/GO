// File name: ...\s02\19_utf8_1\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {

	i := 38

	fmt.Printf("%7b %3o %d %x %s \n", i, i, i, i, string(i))

	i = 40
	fmt.Printf("%7b %3o %d %x %s \n", i, i, i, i, string(i))

	ch := 'ɇ'
	fmt.Printf("\n%T\n", ch)
	fmt.Printf("%7b %3o %d %x %s \n", ch, ch, ch, ch, string(ch))

	ch = 'ξ'
	fmt.Printf("\n%T\n", ch)
	fmt.Printf("%7b %3o %d %x %s \n", ch, ch, ch, ch, string(ch))

	i = 958
	fmt.Printf("%7b %3o %d %x %s \n", i, i, i, i, string(i))

	ch = '本'
	fmt.Printf("\n%T\n", ch)
	fmt.Printf("%7b %3o %d %x %s \n\n", ch, ch, ch, ch, string(ch))

	ch = 'A'
	fmt.Printf("%v \t %v \t\t\t %b\n", string(ch), []byte(string(ch)), ch)

	ch = '£'
	fmt.Printf("%v \t %v \t\t %b\n", string(ch), []byte(string(ch)), ch)

	ch = '㪑'
	fmt.Printf("%v \t %v \t\t %b \n", string(ch), []byte(string(ch)), ch)

	ch = '𠜎'
	fmt.Printf("%v \t %v \t %b \n", string(ch), []byte(string(ch)), ch)
}
