package main

import "fmt"

// Outside Declaration
var job string = "Data Scientist"

// Non Declared Statement
// Only inside fxn
// gender := "Male"

func main() {
	var foo int // Variable Declaration
	fmt.Println("Variable", foo)
	// Check for datatype
	fmt.Printf("%T\n", foo)

	// Declare A Variable & Initialize/Assignment
	var name string = "Jesse"
	fmt.Println("What is your name?", name)

	// Shorthand
	age := 24
	fmt.Println("Age:", age)

	// Usage in Main Fxn
	fmt.Println("Job:", job)

	// fmt.Println(gender)

	// Declare & Assign Value next
	var a int
	a = 24
	fmt.Println(a)

	// Declaration with Specify Data Type
	var address = "Accra"
	fmt.Println("Address:", address)
	fmt.Printf("%T, %v \n", address, address)

}
