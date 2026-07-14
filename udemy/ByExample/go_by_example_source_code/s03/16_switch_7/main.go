// File name: ...\s03\16_switch_7\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	var x interface{}

	// try other values for x: 10, 10.5, true, "hello", mult, no value, and uint(10)
	x = 10

	switch i := x.(type) {
	case int:
		fmt.Printf("%T %v", i, i)
	case float64:
		fmt.Printf("%T %v", i, i)
	case bool, string:
		fmt.Println("type is bool or string")
	case func(int) float64:
		fmt.Printf("%T", i)
	case nil:
		fmt.Println("x is nil")
	default:
		fmt.Println("don't know the type")
	}
}

func mult(i int) float64 {
	return float64(i) * 1.5
}
