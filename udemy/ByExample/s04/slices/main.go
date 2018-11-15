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
	fmt.Println("Append", c)

	// Copy
	arr := []int{1, 2, 3}
	tmp := make([]int, len(arr))
	copy(tmp, arr)
	fmt.Println("Copy", arr)

	// Cut
	a1 := []int{1, 2, 3, 4, 5, 6, 7}
	a2 := append(a1[0:2], a1[4:]...)
	fmt.Println("Cut", a2)

	// Delete
	{
		rand.Seed(time.Now().UnixNano())
		a1 := []int{1, 2, 3, 4, 5, 6, 7}
		c := rand.Intn(len(a1) - 1)
		a2 := append(a1[0:c], a1[c+1:]...)
		fmt.Println("Delete", a2)
	}
	{
		// Delete without preserving order
		a := []int{1, 2, 3, 4, 5, 6, 7}
		i := 2
		a[i] = a[len(a)-1]
		a = a[:len(a)-1]
		fmt.Println("Delete without preserving order", a)
	}

	//Expand https://github.com/golang/go/wiki/SliceTricks
	{
		a := []int{1, 2, 3, 4, 5, 6, 7}
		b = append(a[:2], append(make([]int, 4), a[2:]...)...)
		fmt.Println("Expand", b)
	}

	{
		// Extend
		a := []int{1, 2, 0, 0, 0, 0, 3, 4, 5, 6, 7}
		a = append(a, make([]int, 4)...)
		fmt.Println("Extend", a)
	}

	{
		// Insert
		a := []int{1, 2, 0, 0, 0, 0, 3, 4, 5, 6, 7}
		b := []int{9, 10}
		a = append(a[:3], append(b, a[3:]...)...)
		fmt.Println("Insert", a)
	}

}
