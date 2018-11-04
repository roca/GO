// File name: ...\s04\16_slices_operations\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

// ASSIGNMENT: implementing basic operations with slice.
// Use the slicing techniques you've learned so far to implement
// the following basic operations with slices:
// Append (a=[1 2 3] & b=[4 5 6 7] -> c=[1 2 3 4 5 6 7])
// Copy (a=[1 2 3] -> b=[1 2 3])
// Cut (a=[1 2 3 4 5 6 7] -> a=[1 2 5 6 7])
// Delete (with preserving order)    (a=[1 2 5 6 7] -> a=[1 5 6 7])
// Delete (without preserving order) (a=[1 2 3 4 5 6 7] -> a=[1 2 7 4 5 6])
// Expand (a=[1 2 7 4 5 6] -> a=[1 2 0 0 0 0 7 4 5 6])
// Extend (a=[1 2 0 0 0 0 7 4 5 6] -> a=[1 2 0 0 0 0 7 4 5 6 0 0 0])
// Insert (a=[1 2 0 0 0 0 7 4 5 6 0 0 0] -> a=[1 2 0 9 10 0 0 0 7 4 5 6 0 0 0])
//
// Ref: https://github.com/golang/go/wiki/SliceTricks
//
func main() {

	a := []int{1, 2, 3}
	b := []int{4, 5, 6, 7}

	a = append(a, b...) // Append
	fmt.Printf(" (1) a%v b%v \n", a, b)

	c := make([]int, len(a))
	copy(c, a) // Copy
	fmt.Printf(" (2) a%v c%v \n", a, c)

	d := append([]int(nil), a...) // Copy
	fmt.Printf(" (3) a%v d%v \n", a, d)

	c = append(c[:2], c[4:]...) // Cut
	fmt.Printf(" (4) c%v \n", c)

	c = append(c[:1], c[2:]...) // Delete (with preserving order)
	fmt.Printf(" (5) c%v \n", c)

	c = c[:1+copy(c[1:], c[2:])] // Delete (with preserving order)
	fmt.Printf(" (6) c%v \n", c)

	// Delete (without preserving order)
	fmt.Printf(" (7) d%v \n", d)
	d[2] = d[len(d)-1]
	d = d[:len(d)-1]
	fmt.Printf(" (8) d%v \n", d)

	d = append(d[:2], append(make([]int, 4), d[2:]...)...) // Expand
	fmt.Printf(" (9) d%v \n", d)

	d = append(d, make([]int, 3)...) // Extend
	fmt.Printf("(10) d%v \n", d)

	d = append(d[:3], append([]int{9, 10}, d[3:]...)...) // Insert
	fmt.Printf("(11) d%v \n", d)
}
