// File name: ...\s07\01_pointers_demo\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	x := 1
	p := &x
	q := &x

	fmt.Printf(" x =>  type=%T value=%d  \n", x, x)
	fmt.Printf("&x =>  type=%T value=%x  \n", &x, &x)

	fmt.Println()
	fmt.Printf("*p =>  type=%T value=%d  \n", *p, *p)
	fmt.Printf(" p =>  type=%T value=%x  \n", p, p)
	fmt.Printf("&p =>  type=%T value=%x  \n", &p, &p)

	fmt.Println()
	fmt.Printf("*q =>  type=%T value=%d  \n", *q, *q)
	fmt.Printf(" q =>  type=%T value=%x  \n", q, q)
	fmt.Printf("&q =>  type=%T value=%x  \n", &q, &q)
}
