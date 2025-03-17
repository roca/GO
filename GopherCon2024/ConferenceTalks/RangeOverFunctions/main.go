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
	// Print all elements
	set.PrintAllElemetsPush(u)

	for next, stop := range u.Pull() {
		v, ok := next()
		if !ok {
			break
		}
		fmt.Println(v)
	}
	stop()

}
