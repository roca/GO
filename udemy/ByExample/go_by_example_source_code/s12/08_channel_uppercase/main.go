// File name: ...\s12\08_channel_uppercase\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"strings"
	"time"
)

// ASSIGNMENT - Write a program that contains two channels that
// accept a text and convert the text to uppercase and lowercase.

func main() {
	text := "Please Modify Me!"

	cu := make(chan string) //cu for cUppercase
	cl := make(chan string) //cl for cLowercase

	fmt.Println("(m) before uppercase()")
	go uppercase(text, cu)
	fmt.Println("(m) before lowercase()")
	go lowercase(text, cl)
	fmt.Println("(m) after both")

	sUpper, sLower := <-cu, <-cl
	fmt.Printf("sUpper=%s sLower=%s\n", sUpper, sLower)
}

func uppercase(s string, c chan string) {
	fmt.Println("(f) entering uppercase()")
	time.Sleep(3 * time.Second)

	c <- strings.ToUpper(s)
	fmt.Println("(f) exiting from uppercase(); This line may or may not be reached.")
}

func lowercase(s string, c chan string) {
	fmt.Println("(f) entering lowercase()")
	time.Sleep(3 * time.Second)

	c <- strings.ToLower(s)
	fmt.Println("(f) exting from lowercase(); This line may or may not be reached.")
}
