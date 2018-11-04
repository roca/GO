// File name: ...\s10\06_interface_sort1\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"sort"
)

func main() {

	n := []int{7, 2, 10, 14, 1, 14, 9}
	fmt.Println(n)
	sort.Ints(n)
	fmt.Println(n)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	s := []string{"Susan", "Tyler", "Ava", "Nick"}
	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)
}
