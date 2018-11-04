// File name: ...\s05\08_maps_search\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

// ASSIGNMENT - Create a map for data like "Golang": ".go" and search for some values.
func main() {

	fileExt := map[string]string{
		"Golang": ".go",
		"C++":    ".cpp",
		"Java":   ".java",
		"Python": ".py",
	}

	fmt.Println(fileExt)
	fmt.Println(len(fileExt))

	if ext, ok := fileExt["Java"]; ok {
		fmt.Println(ext, ok)
	}

	if ext, ok := fileExt["Ruby"]; ok {
		fmt.Println(ext, ok)
	} else {
		fmt.Println("Not found!")
	}
}
