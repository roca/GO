package main

import (
	"fmt"
	"math/rand"
	"time"
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

	// Delete
	{
		rand.Seed(time.Now().UnixNano())
		a1 := []int{1, 2, 3, 4, 5, 6, 7}
		c := rand.Intn(len(a1) - 1)
		fmt.Println(c)
		a2 := append(a1[0:c], a1[c+1:]...)
		fmt.Println(a2)
	}

	//Expand https://github.com/golang/go/wiki/SliceTricks
	{
		a := []int{1, 2, 3, 4, 5, 6, 7}
		b = append(a[:2], append(make([]int, 4), a[2:]...)...)
		fmt.Println(b)
	}

}
