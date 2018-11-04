// File name: ...\s06\01_func_hello\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	sayHi()
	sayGreetings("ESP")
	sayGreetings("GER")
	sayGreetings("ENG")

	fmt.Println()
	fmt.Printf("Type of sayHi: %T\n", sayHi)
	fmt.Printf("Type of sayGreetings: %T\n", sayGreetings)
}
func sayHi() {
	fmt.Println("Hi there!")
}

func sayGreetings(lang string) {
	if lang == "ESP" {
		fmt.Println("Hola")
	} else if lang == "GER" {
		fmt.Println("Hallo")
	} else {
		fmt.Println("Hello")
	}

}
