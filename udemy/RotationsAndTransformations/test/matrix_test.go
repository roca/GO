package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"udemy.com/aml/matrix"
)

// Case 01: Default valuses for Matrixshould be 0.0
func TestCase01(t *testing.T) {
	m := matrix.Matrix{}

	expected := []float64{
		0.0, 0.0, 0.0,
		0.0, 0.0, 0.0,
		0.0, 0.0, 0.0,
	}
	actual := []float64{
		m.M11, m.M12, m.M13,
		m.M21, m.M22, m.M23,
		m.M31, m.M32, m.M33,
	}

	assert.Equal(t, actual, expected, "Default valuses for Vector should be 0.0 float64")
}

// func TestMatPrint(t *testing.T) {
// 	v := make([]float64, 12)
// 	for i := 0; i < 12; i++ {
// 		v[i] = float64(i)
// 	}
// 	// Create a new matrix
// 	A := mat.NewDense(3, 4, v)
// 	println("A:")
// 	formatedString := aml.MatPrint(A)

// 	fa := mat.Formatted(A, mat.Prefix(""), mat.Squeeze())

// 	assert.Equal(t, formatedString, fmt.Sprintf("%v\n", fa), "Matrix string output did not match")

// }
