package main

import "fmt"

func main() {

	// Declare An Array
	var arr [2]int
	// Assign Values
	arr[0] = 5

	// Read the Element/Value
	i := arr[0]
	fmt.Println(i)
	fmt.Println("Length", len(arr))
	fmt.Printf("%T \n", arr)

	// Method 2 : Declare & Assign & Initialize with Values
	var even = [4]int{2, 4, 6, 8}
	fmt.Println(even)

	var oddfloat = [4]float64{1, 3.3, 5.3, 7.7}
	fmt.Println(oddfloat)

	// Method 3: Shorthand
	arr2 := [2]int{2, 5}
	fmt.Println("Array 2:", arr2)

	// With Strings
	food := [3]string{"Rice", "Fufu", "Bread"}
	fmt.Println(food)

	fmt.Printf("%T \n", food)

	//  Compilier Specify Size
	lang := [...]string{"Go", "Julia", "Python", "JS"}
	fmt.Printf("%T \n", lang)

}
