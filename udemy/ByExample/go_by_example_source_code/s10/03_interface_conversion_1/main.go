// File name: ...\s10\03_interface_conversion_1\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	var i = 10
	var f = 10.25
	fmt.Println(f + float64(i))
	fmt.Println(int(f) + i)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	var r rune = 'a'
	var i2 int32 = 'b'

	fmt.Println(r, i2)
	fmt.Println(string(r), string(i2))

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	b := []byte{'t', 'e', 's', 't'}
	fmt.Println(b, string(b))

	s := "test"
	fmt.Println([]byte(s), s)
}
