package main

import "fmt"

var i = 100

func main() {
	fmt.Println("Scope global", i)
	fmt.Println("Address global",&i)
	i := 200
	fmt.Println("Scope main", i)
	fmt.Println("Address main", &i)
}
