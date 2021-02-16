package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
	"udemy.com/aml"
)

func TestMatPrint(t *testing.T) {
	v := make([]float64, 12)
	for i := 0; i < 12; i++ {
		v[i] = float64(i)
	}
	// Create a new matrix
	A := mat.NewDense(3, 4, v)
	println("A:")
	formatedString := aml.MatPrint(A)

	fa := mat.Formatted(A, mat.Prefix(""), mat.Squeeze())

	assert.Equal(t, formatedString, fmt.Sprintf("%v\n", fa), "Matrix string output did not match")

}
