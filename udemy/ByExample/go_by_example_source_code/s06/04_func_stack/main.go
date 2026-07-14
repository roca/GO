// File name: ...\s06\04_func_stack\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

var i = 10

func main() {
	f1()
}
func f1() {
	j := 20
	k := 30
	fmt.Println("f1(#1) -", i, j, k)
	f2(j)
	fmt.Println("f1(#2) -", i, j, k)
}

func f2(jp int) {

	// fmt.Println("f2() -", j) //compiler error
	// fmt.Println("f2() -", k) //compiler error

	fmt.Println("f2(#1) -", i, jp)
	i := 30

	f3(i, jp)
	fmt.Println("f2(#2) -", i, jp)
}

func f3(ip, jpp int) {
	ip++
	jpp++
	fmt.Println("f3(#1) -", i, ip, jpp)
}
