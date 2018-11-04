// File name: ...\s08\11_func_defer_2\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

// ASSIGNMENT - Without running the following program find its output.

var i = 0

func main() {
	fmt.Println("i(m1)=", i)
	testDefer()
	fmt.Println("i(m2)=", i)
}

func testDefer() {

	fmt.Println("i(t1)=", i)

	defer closeFiles()
	defer closeDBConnections()

	fmt.Println("i(t2)=", i)
	doSomehting()
	fmt.Println("i(t3)=", i)
}

func closeFiles() {
	fmt.Println("i(f1)=", i)
	i = 1
}

func closeDBConnections() {
	fmt.Println("i(f2)=", i)
	i += 2
}

func doSomehting() {
	fmt.Println("i(f3)=", i)
	i = 3
}
