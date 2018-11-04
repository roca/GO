// File name: ...\s06\08_func_slice_stack\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

// **NOTE** ASSIGNMENT: Use 'slice' and variadic functions to implement a 
// simple stack with three operations: 
// push: adding elements to the top of stack (push is normally adding one 
// element, but in this example we want to use a variadic function);
// pop: removing one element from the top of stack;
// top: retrieving the head element of the stack.

func main() {
	s := []int{}

	s = push(s, 1, 2, 3)
	fmt.Println(s)

	s = pop(s)
	fmt.Println(s)

	s = push(s, 10, 11)
	fmt.Println(s)

	fmt.Println(top(s))
}

func push(s []int, newS ...int) []int {
	return append(s, newS...)
}

func pop(s []int) []int {
	return s[:len(s)-1]
}

func top(s []int) int {
	return s[len(s)-1]
}
