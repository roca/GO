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

func TestPlus(t *testing.T) {
	v1 := &aml.Vector3{1.0,2.0,3.0}
	v2 := &aml.Vector3{1.0,2.0,3.0}

	v1.Plus(v2)
	assert.Equal(t,v1.X, 2.0)
	assert.Equal(t,v1.Y,4.0)
	assert.Equal(t,v1.Z,6.0)
}
