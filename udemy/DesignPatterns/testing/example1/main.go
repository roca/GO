package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])
	result := sum(a, b)
	fmt.Printf("The sum of %d and %d is %d\n", a, b, result)
}

func sum(a, b int) int {
	return a + b
}
