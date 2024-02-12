package main

import (
	"fmt"
	"math"
)

func alphabetSize(numBits int) float64 {
	sum := 1.0
	for i := 0; i < numBits; i++ {
		sum += math.Pow(2.0, float64(i))
	}
	return sum
}

// don't touch below this line

func test(num int) {
	fmt.Printf("Alphabet size for %v bits: %v\n", num, alphabetSize(num))
}

func main() {
	for i := 1; i < 17; i++ {
		test(i)
	}
}
