// File name: ...\s02\25_scope\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

var g int = 10

// g := 20

func main() {

	fmt.Println(g)

	{
		g := 20
		fmt.Println(g)
		f := 2.3
		fmt.Println(f)
	}

	// fmt.Println(f) //compiler error - undefined 'f'

	fmt.Println(g)
	g += 1

	{
		g += 4
		fmt.Println(g)
	}

	sayHello()
	// fmt.Println(s)		//compiler error - undefined 's'
}

func sayHello() {
	fmt.Println("Hello!")
	fmt.Println(g)

	s := "from testScope()"
	fmt.Println(s)
}
