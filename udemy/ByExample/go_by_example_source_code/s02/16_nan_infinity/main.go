// File name: ...\s02\16_nan_infinity\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"math"
)

func main() {

	var x float64
	fmt.Println(x, -x, 1/x, -1/x, x/x)

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)
}
