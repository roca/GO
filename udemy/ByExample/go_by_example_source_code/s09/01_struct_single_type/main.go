// File name: ...\s09\01_struct_single_type\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	type myType float64
	var total myType

	total = 44
	fmt.Printf("%.2f %T \n", total, total)

	var total2 float64

	// total2 = total //compiler error
	total2 = float64(total)
	fmt.Printf("%.2f %T \n", total2, total2)

	// if total == total2 { //compiler error - mismatched types
	// 	fmt.Println("xx")
	// }
}
