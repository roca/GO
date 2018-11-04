// File name: ...\s02\08_string\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	fmt.Println(len("Good Day"))
	fmt.Println("Good Day"[5])
	fmt.Println()

	s := "SuperHelpful"

	fmt.Println(s[:5])
	fmt.Println(s[5:])
	fmt.Println(s[2:9])
	fmt.Println(s[:])

	fmt.Println(s[:5] + s[5:])

	//s[0] = 'D' // compile error:

	s += "Friend"
	fmt.Println(s)
}
