// File name: ...\s05\07_maps_unique\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

// ASSIGNMENT: create a slice of string, including some repeating words.
// Write a program to remove the repeating words.
func main() {
	s := []string{"one", "two", "two", "three", "four", "four", "one", "four"}

	m := make(map[string]bool)

	for _, v := range s {
		m[v] = true
	}

	s = s[:len(m)]
	i := 0
	for key := range m {
		s[i] = key
		i++
	}

	fmt.Println(s)
	fmt.Println(m)
}
