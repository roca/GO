package main

import (
	"gonum.org/v1/gonum/mat"
	"udemy.com/aml"
)

func main() {
	v := make([]float64, 12)
	for i := 0; i < 12; i++ {
		v[i] = float64(i)
	}
	// Create a new matrix
	A := mat.NewDense(3, 4, v)
	println("A:")
	aml.MatPrint(A)
}
