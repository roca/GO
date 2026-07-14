package main

import "fmt"

func main() {

	// Slices
	// Method 1
	// Declare & Initiailize with Values
	var slice = []int{3, 5, 7, 9}
	fmt.Println(slice)
	// Check For Type
	fmt.Printf("%T \n", slice)
	// Length/Size
	fmt.Println("Length", len(slice))
	fmt.Println("Capacity", cap(slice))

	// Append Values
	// s2 := append(slice, 20, 21)
	// fmt.Println(s2)
	// fmt.Println("Length", len(s2))

	// // Method 2:
	// lang := []string{0: "Go", 1: "Julia", 3: "Python"}
	// fmt.Println(lang)

	// Indexing/Slicing
	// a := slice[0:2]
	// a := slice[:] // Prints Everything stored
	// a := slice[1:] // From Index 1 and upwards
	a := slice[:3] // From Index 3 and downwards

	fmt.Println(a)

	// Method 3: make
	// m: slice,map(dictionary,key=value),channel
	// make([]type,len,cap)
	// m := make([]byte, 3, 3)
	slicestring := make([]string, 4, 4)
	fmt.Printf("%T\n", slicestring)
}
