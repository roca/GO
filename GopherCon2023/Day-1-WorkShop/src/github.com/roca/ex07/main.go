package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type intTypes interface {
	//uint | uint8 | uint16 | uint32 | uint64 | uintptr
	constraints.Unsigned
}

func sumSlice[S []E, E intTypes](s S) E {
	var sum E
	for _, v := range s {
		sum += v
	}
	return sum
}

func main() {
	fmt.Printf("%v\n", sumSlice([]uint{1, 2, 3}))

}
