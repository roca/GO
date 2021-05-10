package main

import "fmt"

func main() {

	// Ranges: loop over array/slice
	var evennum = [5]int{2, 4, 6, 8, 10}
	// Slice
	var oddnum = []int{2, 4, 6, 8, 10}

	// index and element
	for index, elem := range evennum {
		fmt.Println(index, "=", elem)
	}

	// without index but only element
	// _ blank identifier/operator
	for _, elem := range evennum {
		fmt.Println(elem + 1)

	}
	// Only Index but without element
	for index := range evennum {
		fmt.Println("Index::", index)
	}

	// without index but only element
	// _ blank identifier/operator
	for _, elem := range oddnum {
		fmt.Println(elem + 1)

	}
}
