// File name: ...\s02\09_binary\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	fmt.Printf("%3d %8b %08b \n", 2, 2, 2)		
	fmt.Printf("%3d %8b %08b \n\n", -2, -2, -2) 

	fmt.Printf("%3d %08b \n", 3, 3)
	fmt.Printf("%3d %08b \n\n", -3, -3)

	fmt.Printf("%3d %08b \n", 10, 10)
	fmt.Printf("%3d %08b \n\n", -10, -10)
}
