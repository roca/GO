// File name: ...\s08\03_func_return_function\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	hi := sayGreetings("ESP")
	fmt.Println(hi())

	hi = sayGreetings("GER")
	fmt.Println(sayGreetings("ENG")())

	fmt.Println()
	fmt.Printf("%T \n", hi)
	fmt.Printf("%T \n", hi())
}

func sayGreetings(lang string) func() string {

	hi := func() string {
		return "Hello"
	}

	if lang == "ESP" {
		hi = func() string {
			return "Hola"
		}
	} else if lang == "GER" {
		hi = func() string {
			return "Hallo"
		}
	}

	return hi
}
