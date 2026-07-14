// File name: ...\09_maps_composite_literal\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

// Ref: https://golang.org/doc/effective_go.html?#maps
func main() {
	timeZone := map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}

	if offset, ok := timeZone["EST"]; ok {
		fmt.Println(offset)
	}
}
