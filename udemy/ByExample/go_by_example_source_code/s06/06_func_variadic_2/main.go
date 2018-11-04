// File name: ...\s06\06_func_variadic_2\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

// ASSIGNMENT - Write a program to pass a slice of string to a 
// function (variadic) to loop through the elements and print them.
func main() {
	names := []string{"Amy", "Rob", "Helen"}
	echo(names...)				//var...
}

func echo(names ...string) {	//...type
	for _, s := range names {
		fmt.Print(s, " ")
	}
}
