package main

import "fmt"

type Salutation struct {
	name     string
	greeting *string
}

const (
	PI       = 3.14
	Language = "Go"
)

const (
	A = iota
	B
	C
)

func main() {

	var message Salutation
	message.name = "Hello Go World"
	message.greeting = &message.name
	*message.greeting = "hi"

	fmt.Println(message.name, *message.greeting)

	fmt.Println(A, B, C)
	fmt.Println(PI)
	fmt.Println(Language)
}
