package main

import "fmt"

func main() {
	var a any

	a = 7
	fmt.Println("a:", a)

	a = "Hi"
	fmt.Println("a:", a)

	s := a.(string)
	fmt.Println("s:", s)

	i, ok := a.(int)
	if ok {
		fmt.Println("i:", i)
	} else {
		fmt.Printf("not an int (%T)\n", a)
	}
}
