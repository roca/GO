// File name: ...\s07\09_pointers_by_ref2
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	a := 1
	fmt.Printf("(1) a=%d &a=%x \n\n", a, &a)

	f3(&a)
	fmt.Printf("(2) a=%d \n\n", a)

	fmt.Printf("(3) a=%d\n", f3(&a))

}

func f3(y *int) int {
	*y++
	fmt.Printf("(f) *y=%d y=%x\n", *y, y)
	return *y
}
