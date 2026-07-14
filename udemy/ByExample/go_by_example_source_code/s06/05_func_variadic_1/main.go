// File name: ...\s06\05_func_variadic_1\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	sum(1, 2)
	sum(3, 4, 5)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)

	sum2(nums)

	fmt.Printf("sum=%T sum2=%T \n", sum, sum2)
}

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Printf("nums=%v total=%d type=%T\n", nums, total, nums)
}

func sum2(nums []int) {
	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Printf("nums=%v total=%d type=%T\n", nums, total, nums)
}
