package main

import (
	"fmt"
	"math"
)

func Cbrt(x complex128) complex128 {
	isGoodEnough := func(guess complex128) bool {
		z := (guess * guess * guess) - x
		e := math.Sqrt((real(z) * real(z)) + (imag(z) * imag(z)))
		return e < 0.0001
	}
	improve := func(guess complex128) complex128 {
		return guess - (((guess * guess * guess) - x) / (3 * guess * guess))
	}
	cbrtItr := func(guess complex128) complex128 { return complex(0, 0) }
	cbrtItr = func(guess complex128) complex128 {

		if isGoodEnough(guess) {
			return guess
		} else {
			return cbrtItr(improve(guess))
		}
	}
	return cbrtItr(complex(1.0, 1.0))

}

func main() {
	fmt.Println(Cbrt(2))
}
