// tour58
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return "cannot Sqrt negative number: "
}

func precision(x, y float64) float64 {
	p := 1.0 - (math.Abs(x-y) / x)
	return (100.0 * p)

}

func Sqrt(x float64) float64 {
	var e ErrNegativeSqrt
	e = x
	if x < 0 {
		return e
	}
	z := float64(1)
	i := 0

	for {
		i++
		znew := z - (((z * z) - x) / (2 * z))
		if precision(z, znew) > 99.99 {
			break
		}
		z = znew
		//fmt.Println("iterartion ", i, " : ", z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
