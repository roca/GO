// File name: ...\s07\08_pointers_by_ref1\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	x := 10

	fmt.Printf("x=%d\n", x)
	f2(&x)
	fmt.Printf("x=%d\n", x)
}

func f2(y *int) {
	fmt.Printf("(f2) y=%d\n", *y)
	*y += 2
	fmt.Printf("(f2) y=%d\n", *y)
}
