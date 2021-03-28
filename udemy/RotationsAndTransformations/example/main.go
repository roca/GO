package main

import (
	"fmt"

	"udemy.com/aml/dcm"
	"udemy.com/aml/matrix"
	"udemy.com/aml/vector"
)

func main() {
	R := matrix.Identity()
	for i := 0; i < 100; i++ {
		v, _ := vector.New([]float64{1., 0.0, 0.0})
		RDot, _ := dcm.KinematicRatesFromBodyRates(&R, &v)
		R, _ = dcm.Intergrate(&R, &RDot, .01)

		fmt.Printf("%f\n", R.Data())
	}
}
