package main

import (
	"./greeting"
)

func main() {

	var s = greeting.Salutation{"Bob", "Hello"}
	greeting.Greet(s, greeting.CreatePrintFunction("!!!"), false)
}
