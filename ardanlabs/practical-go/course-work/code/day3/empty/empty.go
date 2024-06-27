package main

import "fmt"

func main() {
	var i any

	i = 7
	fmt.Println(i)

	i = "Hello"
	fmt.Println(i)

	// Rule of thumb: Don't use empty interface :)


	s := i.(string) // type assertion
	fmt.Println(s)
	fmt.Printf("%#v\n", i)

	// n := i.(int) // panic: interface conversion: interface {} is string, not int
	// fmt.Println(n)

	n,ok := i.(int)
	if ok {
		fmt.Println(n)
	} else {
		fmt.Printf("'%v' is not an int",i)
	}

}
