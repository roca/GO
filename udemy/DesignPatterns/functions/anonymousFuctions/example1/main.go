package main

import "fmt"


func main() {
    addN := func(m int) (func(int)int){
		// Returns a function with a closure around m
		return func(n int) int{
			return m + n
		}
	}

	addFiveTo := addN(5) // Closure around 5. 
	result := addFiveTo(6)
	// 5 + 6 must print 11
	fmt.Println(result)
}