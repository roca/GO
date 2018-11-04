// File name: ...\s04\06_arrays_multi_dim\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {
	a := [2][3]int{{2, 4, 6}, {4, 6, 8}}
	fmt.Printf("a=%v type=%T, len=%d\n", a, a, len(a))

	var b [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			b[i][j] = (i + j + 1) * 2
		}
	}
	fmt.Printf("b=%v type=%T, len=%d\n", b, b, len(b))

	a[0][1] = 5
	a[1][1] = 10
	fmt.Printf("a=%v type=%T, len=%d\n", a, a, len(a))
}
