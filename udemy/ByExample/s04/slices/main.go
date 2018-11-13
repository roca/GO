package main

import (
	"fmt"
)

func main() {

	// Append
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := append(a, b...)
	fmt.Println(c)

	// Copy
	arr := []int{1, 2, 3}
	tmp := make([]int, len(arr))
	copy(tmp, arr)
	fmt.Println(tmp)
	fmt.Println(arr)

	// Cut
	a1 := []int{1, 2, 3, 4, 5, 6, 7}
	a2 := append(a1[0:2], a1[4:]...)
	fmt.Println(a2)

}
