package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"udemy.com/aml/vector"
)

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

// Case 03: Should initialize with slice of three values
func TestCase04(t *testing.T) {
	v, _ := vector.New([]float64{3.0, 2.0, 1.0})

	expected := []float64{3.0, 2.0, 1.0}
	actual := []float64{v.X, v.Y, v.Z}

	assert.Equal(t, actual, expected, "Should initialize with slice of three values")
}
