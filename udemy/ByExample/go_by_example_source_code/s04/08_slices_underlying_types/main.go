// File name: ...\s04\08_slices_underlying_types
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {
	var a [3]int
	fmt.Printf("a=%T \n\n", a)

	b := [...]int{1, 2, 3, 4, 5}
	b1 := b[:3]
	fmt.Printf("b=%T b1=%T\n\n", b, b1)

	c := [5]int{1, 2, 3, 4, 5}
	c1 := c[:3]
	fmt.Printf("c=%T c1=%T\n\n", c, c1)

	d := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	d1 := d[:3]
	fmt.Printf("d=%T d1=%T\n\n", d, d1)
}
