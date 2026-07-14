// File name: ...\s04\12_slices_make\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {
	nums := make([]int, 2, 3) // or // nums := new([3]int)[0:2]

	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
	nums[0] = 1
	nums[1] = 3
	//nums[2] = 3	//panic: runtime error: index out of range
	fmt.Printf("%v\n\n", nums)

	s1 := []int{1, 3, 5} //1st slice
	s2 := make([]int, 2) //2nd slice

	copy(s2, s1)
	s3 := append(s2, 4, 6) //3rd slice

	fmt.Println(s1, s2, s3)
	fmt.Println(len(s1), len(s2), len(s3))
	fmt.Println(cap(s1), cap(s2), cap(s3))

	fmt.Println()
	for i := range s3 {
		fmt.Print(s3[i], " ")
	}

	fmt.Println()
	for _, el := range s3 {
		fmt.Printf("%d ", el)
	}
}
