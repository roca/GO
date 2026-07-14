package main

import "fmt"

func main() {
	var i any

	i = 7
	fmt.Println(i)

	i = "Hello"
	fmt.Println(i)

	// Rule of thumb: Don't use empty interface :)

	s := i.(string) // type assertion if underlying type is not string, it will panic
	fmt.Println(s)
	fmt.Printf("%#v\n", i)

	// n := i.(int) // panic: interface conversion: interface {} is string, not int
	// fmt.Println(n)

	n, ok := i.(int)
	if ok {
		fmt.Println(n)
	} else {
		fmt.Printf("'%v' is not an int\n", i)
	}

	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Printf("%T i is another type\n", i)
	}

	// fmt.Println(maxInts([]int{8, 2, 3, 4, 5}))
	// fmt.Println(maxFloat64s([]float64{5.7, 2.2, 3.3, 4.4, 5.5}))

	fmt.Println(max([]int{8, 2, 3, 4, 5}))               // 8
	fmt.Println(max([]float64{5.7, 2.2, 3.3, 4.4, 5.5})) // 5.7

}

type Number interface {
	int | float64
}

// func max[T int | float64](nums []T) T {
func max[T Number](nums []T) T {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

func maxInts(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

func maxFloat64s(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}
