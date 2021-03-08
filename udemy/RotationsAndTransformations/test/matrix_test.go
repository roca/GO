package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"udemy.com/aml/matrix"
	"udemy.com/aml/vector"
)

// Case 01: Default valuses for Matrixshould be 0.0
func TestConstructWithNoData(t *testing.T) {
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

func TestConstructWithSingelScalar(t *testing.T) {
	s := 5.
	m, _ := matrix.New(s)

	expected := []float64{
		s, s, s,
		s, s, s,
		s, s, s,
	}
	actual := []float64{
		m.M11, m.M12, m.M13,
		m.M21, m.M22, m.M23,
		m.M31, m.M32, m.M33,
	}
	assert.Equal(t, actual, expected, "Valuses for Matrix should be %f ", s)
}

func TestConstructWithNineScalars(t *testing.T) {
	m, _ := matrix.New(1., 2., 3., 4., 5., 6., 7., 8., 9.)

	expected := []float64{
		1.0, 2.0, 3.0,
		4.0, 5.0, 6.0,
		7.0, 8.0, 9.0,
	}
	actual := []float64{
		m.M11, m.M12, m.M13,
		m.M21, m.M22, m.M23,
		m.M31, m.M32, m.M33,
	}
	assert.Equal(t, actual, expected, "Valuses for Matrix for each element incorrectly assigned")
}

func TestConstructWithSlice(t *testing.T) {
	m, _ := matrix.New([]float64{1., 2., 3., 4., 5., 6., 7., 8., 9.})

	expected := []float64{
		1.0, 2.0, 3.0,
		4.0, 5.0, 6.0,
		7.0, 8.0, 9.0,
	}
	actual := []float64{
		m.M11, m.M12, m.M13,
		m.M21, m.M22, m.M23,
		m.M31, m.M32, m.M33,
	}
	assert.Equal(t, actual, expected, "Valuses for Matrix for each element incorrectly assigned")
}

func TestConstructWith2DSlice(t *testing.T) {
	m, _ := matrix.New([][]float64{
		{1., 2., 3.},
		{4., 5., 6.},
		{7., 8., 9.},
	})

	expected := []float64{
		1.0, 2.0, 3.0,
		4.0, 5.0, 6.0,
		7.0, 8.0, 9.0,
	}
	actual := []float64{
		m.M11, m.M12, m.M13,
		m.M21, m.M22, m.M23,
		m.M31, m.M32, m.M33,
	}
	assert.Equal(t, actual, expected, "Valuses for Matrix for each element incorrectly assigned")
}

func TestConstructWithVectors(t *testing.T) {
	v1,_ := vector.New(1., 2., 3.)
	v2,_ := vector.New(4., 5., 6.)
	v3,_ := vector.New(7., 8., 9.)
	m, _ := matrix.New(v1, v2, v3)

	expected := []float64{
		1.0, 2.0, 3.0,
		4.0, 5.0, 6.0,
		7.0, 8.0, 9.0,
	}
	actual := []float64{
		m.M11, m.M12, m.M13,
		m.M21, m.M22, m.M23,
		m.M31, m.M32, m.M33,
	}
	assert.Equal(t, actual, expected, "Valuses for Matrix for each element incorrectly assigned")
}
func TestConstructWithSliceOfVectors(t *testing.T) {
	v1,_ := vector.New(1., 2., 3.)
	v2,_ := vector.New(4., 5., 6.)
	v3,_ := vector.New(7., 8., 9.)
	m, _ := matrix.New([]vector.Vector{v1, v2, v3})

	expected := []float64{
		1.0, 2.0, 3.0,
		4.0, 5.0, 6.0,
		7.0, 8.0, 9.0,
	}
	actual := []float64{
		m.M11, m.M12, m.M13,
		m.M21, m.M22, m.M23,
		m.M31, m.M32, m.M33,
	}
	assert.Equal(t, actual, expected, "Valuses for Matrix for each element incorrectly assigned")
}