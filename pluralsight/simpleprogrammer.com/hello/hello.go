package main

import "fmt"

type Salutation struct {
	name     string
	greeting *string
}

func main() {

	var message Salutation
	message.name = "Hello Go World"
	message.greeting = &message.name
	*message.greeting = "hi"

	fmt.Println(message.name, *message.greeting)
}
