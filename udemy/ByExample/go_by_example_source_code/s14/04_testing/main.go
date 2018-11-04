// File name: ...\s14\04_testing\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"

	"./myutil"
)

func main() {
	testCase1()
	fmt.Println()
	testCase2()
}
func testCase1() {
	fmt.Printf("(p) Rectangle Area= %d\n", myutil.Area(2, 3))   //6
	fmt.Printf("(p) Rectangle Perim= %d\n", myutil.Perim(2, 3)) //10

	fmt.Printf("(p) Rectangle Area= %d\n", myutil.Area(5, 10))   //50
	fmt.Printf("(p) Rectangle Perim= %d\n", myutil.Perim(5, 10)) //30

	fmt.Printf("(p) Rectangle Area= %d\n", myutil.Area(4, 2))   //8
	fmt.Printf("(p) Rectangle Perim= %d\n", myutil.Perim(4, 2)) //12
}

func testCase2() {
	fmt.Printf("(p) Average= %.2f\n", myutil.Avg([]float32{4, 6}))            //5.00
	fmt.Printf("(p) Average= %.2f\n", myutil.Avg([]float32{3, -3}))           //0.00
	fmt.Printf("(p) Average= %.2f\n", myutil.Avg([]float32{3, -3, 5, -5, 4})) //0.80
	fmt.Printf("(p) Average= %.2f\n", myutil.Avg([]float32{1, 1, 1, 100}))    //25.75
}
