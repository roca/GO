// File name: ...\s05\10_maps_unicode\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"unicode/utf8"
)

// ASSIGNMENT - In the following text, count the number of distinct Unicode code point.
// "How 你 cómo お元 Wie geht 잘 How 你 có お元 Wie geht 잘 Ho có Wie"
// map is a good data structure to address this assignment.
func main() {
	s := "How 你 cómo お元 Wie geht 잘 How 你 có お元 Wie geht 잘 Ho có Wie"

	r := []rune(s)

	// fmt.Println(r)
	fmt.Println(len(s), utf8.RuneCountInString(s))

	counts := make(map[rune]int)

	fmt.Println()
	for i := range r {
		counts[r[i]]++
	}

	fmt.Printf("%v\n", counts)

	// for k, v := range counts {
	// 	fmt.Printf("%7d %3s %3d  \n", k, string(k), v)
	// }
}
