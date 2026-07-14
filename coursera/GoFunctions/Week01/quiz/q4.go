package main

import "fmt"

func fA() func(string) int {
	return func(s string) int {
		return 1
	}
}

func main() {
	fmt.Println(fA()("x)"))
}
