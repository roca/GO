package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
	"udemy.com/aml"
	"udemy.com/aml/vector"
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

// Vectot test

// Case 01: Default valuses for Vector should be 0.0
func TestCase01(t *testing.T) {
	v := vector.Vector{}

	expected := []float64{0.0, 0.0, 0.0}
	actual := []float64{v.X, v.Y, v.Z}

	assert.Equal(t, actual, expected, "Default valuses for Vector should be 0.0 float64")
}

// Case 02: Should initialize with scalar value
func TestCase02(t *testing.T) {
	v, _ := vector.New(3.0)

	expected := []float64{3.0, 3.0, 3.0}
	actual := []float64{v.X, v.Y, v.Z}

	assert.Equal(t, actual, expected, "Should initialize with scalar value")
}

// Case 03: Should initialize with three values
func TestCase03(t *testing.T) {
	v, _ := vector.New(1.0, 2.0, 3.0)

	expected := []float64{1.0, 2.0, 3.0}
	actual := []float64{v.X, v.Y, v.Z}

	assert.Equal(t, actual, expected, "Should initialize with three values")
}
