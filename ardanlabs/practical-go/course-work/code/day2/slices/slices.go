package main

import (
	"fmt"
	"sort"
)

func main() {
	var s []int                // s is a slice of int
	fmt.Println("len", len(s)) // len is "nil safe"
	if s == nil {
		fmt.Println("s is nil")
	}

	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	sOriginal := make([]int, len(s2)) // make a slice of int with length 10
	copy(sOriginal, s2)
	fmt.Printf("s2 = %#v\n", s2)

	s3 := s2[1:4] // slicing operation, half open range
	fmt.Printf("s3 = %#v\n", s3)

	// fmt.Println(s2[:100]) // panic if out of range
	s3 = append(s3, 100)
	fmt.Printf("s3 (appended) = %#v\n", s3)
	fmt.Printf("s2 (appended) = %#v\n", s2)       // s2 is modified as well!
	fmt.Printf("s (original) = %#v\n", sOriginal) // s2 is modified as well!
	fmt.Printf("s2: len=%d, cap=%d\n", len(s2), cap(s2))
	fmt.Printf("s3: len=%d, cap=%d\n", len(s3), cap(s3))

	var s4 []int
	// var s4 = make([]int, 0, 1_000) // Single allocation
	for i := 0; i < 1_000; i++ {
		s4 = appendInt(s4, i)
	}
	fmt.Printf("s4: len=%d, cap=%d\n", len(s4), cap(s4))
	// s4[1001] = 7 // panic if out of range
	s4 = append(s4, 7)
	fmt.Printf("s4: len=%d, cap=%d\n", len(s4), cap(s4))

	fmt.Println((concat([]string{"A", "B"}, []string{"C", "D", "E"}))) // ["A", "B", "C", "D", "E"]

	v1 := []float64{5, 3, 4, 1, 2}
	v2 := []float64{1, 6, 3, 4, 5, 2}
	fmt.Println(median(v1)) // 3

	m ,_ := median(v2)
	fmt.Println(m) // 3.5

	fmt.Println(v1)
	fmt.Println(v2)

	_, err := median(nil)
	if err != nil {
		fmt.Println(err)
	}
}

func median(values []float64) (float64, error) {
	if values == nil || len(values) == 0 {
		return 0, fmt.Errorf("nil or empty slice given")
	}
	// Copy in order not to modify the original slice
	nums := make([]float64, len(values))
	copy(nums, values)

	sort.Float64s(nums)
	i := len(nums) / 2
	if len(nums)%2 == 0 {
		const n = 2                      // type undefined until when used
		return (nums[i-1] + nums[i]) / n, nil // n is used here and is now a float64
	}
	return nums[i], nil
}

func concat(s1, s2 []string) []string {
	// Restriction: No "for" loop, no "range
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)
	return s

	// return append(s1, s2...)
}

func appendInt(s []int, v int) []int {
	i := len(s)
	if len(s) < cap(s) { // enough space in underlying array
		s = s[:len(s)+1]
	} else { // need to re-allocate and copy
		fmt.Printf("re-allocate: %d->%d\n", len(s), 2*len(s)+1)
		s2 := make([]int, 2*len(s)+1)
		copy(s2, s)
		s = s2[:len(s)+1]
	}

	s[i] = v
	return s
}
