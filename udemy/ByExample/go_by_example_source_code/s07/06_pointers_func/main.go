// File name: ...\s07\06_pointers_func\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	var p1 = f(2)
	fmt.Printf("p1=%x *p1=%d &p1=%x\n\n", p1, *p1, &p1)

	var p2 = f(3)
	fmt.Printf("p2=%x *p2=%d &p2=%x\n", p2, *p2, &p2)
}

func f(inp int) *int {
	v := inp * 2
	fmt.Printf("&v=%x\n", &v)

	return &v
}
