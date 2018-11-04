// File name: ...\s02\12_equality_symbol\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	s := "hi"
	s = "hello"
	
	fmt.Println(s == "hello")
	fmt.Println(s != "hello")

	var (
		name1 = "Tim"
		name2 = "Ava"
		name3 = "Nick"
	)

	fmt.Println(name1, name2, name3)
}
