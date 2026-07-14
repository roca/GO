// File name: ...\s02\02_types_variables\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {

	var i int
	fmt.Println("i =", i)

	j := 2
	fmt.Printf("%d - %T\n", j, j)

	var k1 uint8 = 20
	// k1 = 256		// compiler error - overflow
	// k1 = -2		// compiler error - overflow
	fmt.Printf("%d - %T\n", k1, k1)

	var k2 uint16 = 30

	// k1 = k2			// compiler error - cannot use k2 (type uint16) as type uint8 in assignment
	// k2 = k1			// compiler error
	k2 = uint16(k1)

	fmt.Printf("%d - %T\n", k2, k2)

	fmt.Println(uint64(321325 * 424521))
	// fmt.Println(uint32(321325 * 424521)) //compiler error - overflow
}
