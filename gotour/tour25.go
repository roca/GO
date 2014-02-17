package main

import (
	"fmt"
	"math"
)

func precision(x, y float64) float64 {
	p := 1.0 - (math.Abs(x-y) / x)
	return (100.0 * p)

}

func Sqrt(x float64) float64 {
	z := float64(1)
	i := 0

	for {
		i++
		znew := z - (((z * z) - x) / (2 * z))
		if precision(z, znew) > 99.99 {
			break
		}
		z = znew
		fmt.Println("iterartion ", i, " : ", z)
	}

	return z
}

func main() {
	mySqrt := Sqrt(2)
	mathSqrt := math.Sqrt(2)
	fmt.Println("MySqrt : ", mySqrt)
	fmt.Println("math.Sqrt : ", mathSqrt)
	fmt.Println(precision(mathSqrt, mySqrt), "%")
}
