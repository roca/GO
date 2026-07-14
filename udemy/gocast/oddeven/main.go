package main

import (
	"fmt"
)

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 9, 10}

	for _, value := range ints {
		if value%2 == 0 {
			fmt.Printf("%v is even\n", value)
		} else {
			fmt.Printf("%v is odd\n", value)
		}
	}

}
