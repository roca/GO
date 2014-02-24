// tour58
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %4.2f", e)
}

func precision(x, y float64) float64 {
	p := 1.0 - (math.Abs(x-y) / x)
	return (100.0 * p)

}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
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

	return z, nil
}

func main() {

	answer, err := Sqrt(2)
	if err == nil {
		fmt.Println(answer)
	} else {
		fmt.Println(err)
	}

	answer, err = Sqrt(-2)
	if err == nil {
		fmt.Println(answer)
	} else {
		fmt.Println(err)
	}
}
