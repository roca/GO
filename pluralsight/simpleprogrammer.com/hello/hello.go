package main

import (
	"./greeting"
)

func main() {

	slice := []greeting.Salutation{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "What is up?"},
	}

	greeting.Greet(slice, greeting.CreatePrintFunction("?"), true, 6)
}
