// File name: ...\s07\04_pointers_slice_string\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {

	x := [...]string{"Angela", "Tim", "Nicole", "Henry"}
	p := &x

	y := x[1:3]
	q := &y

	fmt.Printf("*p=%s &p=%x\n", *p, &p)
	fmt.Printf("*q=%s &q=%x\n", *q, &q)

	(*p)[1] = "TIM"
	(*q)[1] = "NICOLE"

	fmt.Println()
	for i := range x {
		fmt.Println(i, x[i], (*p)[i])
	}

	fmt.Println()
	for i := range y {
		fmt.Println(i, y[i], (*q)[i])
	}
}
