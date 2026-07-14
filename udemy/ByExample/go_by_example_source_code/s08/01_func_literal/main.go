// File name: ...\s08\01_func_literal\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

ffunc main() {
	printMsg("Calling a function!")

	showMsg := func(msg string) {
		fmt.Println(msg)
	}

	showMsg("Using a function literal!")
	fmt.Printf("%T\n", showMsg)

	func(msg string) {
		fmt.Println(msg)
	}("quickly reacting!")
}

func printMsg(msg string) {
	fmt.Println(msg)
}
