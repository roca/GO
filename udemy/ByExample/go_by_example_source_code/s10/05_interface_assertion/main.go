// File name: ...\s10\05_interface_assertion\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {

	var message interface{} = "Hello"
	// var message interface{} = 10

	s, ok := message.(string)
	if ok {
		fmt.Printf("%q %T\n", s, message)
	} else {
		fmt.Printf("value is not a string - %q %T\n", s, message)
	}

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	f := 20.355
	fmt.Printf("%d %T %T \n", int(f), int(f), f)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	var num interface{} = 10
	println(num.(int) + 20)
	fmt.Printf("%d %T \n", num, num)
}
