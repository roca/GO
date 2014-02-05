package main

import (
	"./greeting"
)

func main() {

	//greeting.Greet(greeting.Salutation{"Bob", "Hello"}, greeting.CreatePrintFunction("?"), true)
	//greeting.Greet(greeting.Salutation{"Joe", "Hello"}, greeting.CreatePrintFunction("?"), true)
	//greeting.Greet(greeting.Salutation{"Mary", "Hello"}, greeting.CreatePrintFunction("?"), true)
	//greeting.Greet(greeting.Salutation{"John", "Hello"}, greeting.CreatePrintFunction("?"), true)
	//greeting.Greet(greeting.Salutation{"Amy", "Hello"}, greeting.CreatePrintFunction("?"), true)
	//greeting.Greet(greeting.Salutation{"1234567890", "Hello"}, greeting.CreatePrintFunction("?"), true)
	//greeting.Greet(greeting.Salutation{"12345678901", "Hello"}, greeting.CreatePrintFunction("?"), true)
	greeting.TypeSwitchTest(1.345)
	greeting.TypeSwitchTest(1)
	greeting.TypeSwitchTest("hello")
	greeting.TypeSwitchTest(greeting.Salutation{"12345678901", "Hello"})
}
