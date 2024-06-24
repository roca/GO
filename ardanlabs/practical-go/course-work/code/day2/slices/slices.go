package main

import "fmt"

func main() {
	var s []int // s is a slice of int
	fmt.Println("len", len(s)) // len is "nil safe"
	if s == nil {
		fmt.Println("s is nil")
	}

	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("s2 = %#v\n", s2)

	s3 := s2[1:4] // slicing operation, half open range
	fmt.Printf("s3 = %#v\n", s3)

	// fmt.Println(s2[:100]) // panic if out of range
	s3 = append(s3, 100)
	fmt.Printf("s3 (appended) = %#v\n", s3)
	fmt.Printf("s2 (appended) = %#v\n", s2) // s2 is modified as well!
}
