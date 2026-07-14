package main

import (
	"example/set"
	"fmt"
)

func main() {
	// Create a new set
	s := set.New[int]()
	// Add elements
	s.Add(1)
	s.Add(2)
	s.Add(3)
	// Create another set
	t := set.New[int]()
	// Add elements
	t.Add(3)
	t.Add(4)
	t.Add(5)
	// Union
	u := set.Union(s, t)
	fmt.Println("Print all elements using Push")
	set.PrintAllElemetsPush(u)

	fmt.Println("Print all elements using Pull")
	set.PrintAllElementsPull(u)

	fmt.Println("Print all elements using All iterator")
	set.PrintAllElements(u)

	fmt.Println("Print odd elements using All iterator")
	set.PrintOddElements(u)

}