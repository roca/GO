// File name: ...\s06\07_func_slice_del_empty\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

// ASSIGNMENT: Write a function to accept a slice of
// string, and return a new slice with no blank strings.
// Use both variadic and regular functions.
func main() {

	data := []string{"Daisy", "Rose", "", "Tulip"}

	fmt.Printf("%q\n", trimSlice(data)) //form 1
	// fmt.Printf("%q\n", trimSlice(data...)) //form 2

	fmt.Printf("%q\n", data)
}

func trimSlice(data []string) []string { //form1
// func trimSlice(data ...string) []string { //form2

	// newData := data[:0]		//note
	var newData []string
	
	for _, d := range data {
		if d != "" {
			newData = append(newData, d)
		}
	}
	return newData
}

