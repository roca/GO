package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"udemy.com/aml/matrix"
	"udemy.com/aml/vector"
)

func flat(m matrix.Matrix) []float64 {
	return []float64{
		m.M11, m.M12, m.M13,
		m.M21, m.M22, m.M23,
		m.M31, m.M32, m.M33,
	}
}

// Default values for Matrix should be 0.0
func TestConstructWithNoData(t *testing.T) {
	m := matrix.Matrix{}
	expected := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	assert.Equal(t, expected, flat(m), "Default values for Matrix should be 0.0")
}

func TestConstructWithSplat(t *testing.T) {
	s := 5.0
	m := matrix.Splat(s)
	expected := []float64{s, s, s, s, s, s, s, s, s}
	assert.Equal(t, expected, flat(m), "Splat should set every element")
}

func TestConstructWithRows(t *testing.T) {
	m := matrix.New([3][3]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	expected := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, expected, flat(m), "New should assign each element by row")
}

func TestConstructFromRows(t *testing.T) {
	v1 := vector.New(1, 2, 3)
	v2 := vector.New(4, 5, 6)
	v3 := vector.New(7, 8, 9)
	m := matrix.FromRows(v1, v2, v3)
	expected := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, expected, flat(m), "FromRows should assign each row vector")
}

// Value semantics: operations return new matrices, receiver unchanged.
func TestAdditionWithMatrix(t *testing.T) {
	m1 := matrix.New([3][3]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	m2 := matrix.New([3][3]float64{{1.5, 2.5, 3.5}, {4.5, 5.5, 6.5}, {7.5, 8.5, 9.5}})
	m3 := m1.Add(m2)
	expected := [3][3]float64{
		{2.5, 4.5, 6.5},
		{8.5, 10.5, 12.5},
		{14.5, 16.5, 18.5},
	}
	assert.Equal(t, expected, m3.Data(), "Matrix addition values incorrect")
	assert.Equal(t, 1.0, m1.M11, "Add must not mutate the receiver")
}

func TestSubtractionWithMatrix(t *testing.T) {
	m1 := matrix.New([3][3]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	m2 := matrix.New([3][3]float64{{1.5, 2.5, 3.5}, {4.5, 5.5, 6.5}, {7.5, 8.5, 9.5}})
	m3 := m1.Sub(m2)
	expected := [3][3]float64{
		{-0.5, -0.5, -0.5},
		{-0.5, -0.5, -0.5},
		{-0.5, -0.5, -0.5},
	}
	got := m3.Data()
	for i := range 3 {
		assert.InDeltaSlice(t, expected[i][:], got[i][:], 1e-15, "Matrix subtraction values incorrect")
	}
}

func TestMultiplicationWithMatrix(t *testing.T) {
	m1 := matrix.New([3][3]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	m2 := matrix.New([3][3]float64{{1.5, 2.5, 3.5}, {4.5, 5.5, 6.5}, {7.5, 8.5, 9.5}})
	m3 := m1.Mul(m2)
	expected := [3][3]float64{
		{33.0, 39.0, 45.0},
		{73.5, 88.5, 103.5},
		{114.0, 138.0, 162.0},
	}
	got := m3.Data()
	for i := range 3 {
		assert.InDeltaSlice(t, expected[i][:], got[i][:], 1e-13, "Matrix multiplication values incorrect")
	}
}

func TestMultiplicationWithVector(t *testing.T) {
	m := matrix.New([3][3]float64{
		{-2, -3, 2},
		{1, 0, 1},
		{6, -8, 7},
	})
	v := vector.New(-2, 2, 3)
	got := m.MulVec(v)
	expected := []float64{4, 1, -7}
	assert.Equal(t, expected, []float64{got.X, got.Y, got.Z}, "M * V operation is incorrect")
}

func TestScale(t *testing.T) {
	m := matrix.New([3][3]float64{{-2, -3, 2}, {1, -1, 1}, {6, -8, 7}})
	u := m.Scale(0.5)
	expected := [3][3]float64{
		{-1.0, -1.5, 1.0},
		{0.5, -0.5, 0.5},
		{3.0, -4.0, 3.5},
	}
	assert.Equal(t, expected, u.Data(), "Matrix Scale is incorrect")
	assert.Equal(t, -2.0, m.M11, "Scale must not mutate the receiver")
}

func TestDiag(t *testing.T) {
	v := vector.New(-2, -3, 2)
	m := matrix.Diag(v)
	expected := [3][3]float64{
		{-2, 0, 0},
		{0, -3, 0},
		{0, 0, 2},
	}
	assert.Equal(t, expected, m.Data(), "Diag should place the vector on the diagonal")

	u := matrix.New([3][3]float64{{-2, -3, 2}, {1, -1, 1}, {6, -8, 7}})
	d := u.DiagV()
	assert.Equal(t, []float64{-2, -1, 7}, []float64{d.X, d.Y, d.Z}, "DiagV should return the diagonal")
}

func TestTranspose(t *testing.T) {
	m := matrix.New([3][3]float64{{-2, -3, 2}, {1, -1, 1}, {6, -8, 7}})
	u := m.Transpose()
	expected := [3][3]float64{
		{-2, 1, 6},
		{-3, -1, -8},
		{2, 1, 7},
	}
	assert.Equal(t, expected, u.Data(), "Transpose of this matrix is incorrect")
}

func TestDeterminant(t *testing.T) {
	m := matrix.New([3][3]float64{{-2, -3, 2}, {1, -1, 1}, {6, -8, 7}})
	assert.Equal(t, -3.0, m.Determinant(), "Determinant is incorrect")
}

func TestInverse(t *testing.T) {
	m := matrix.New([3][3]float64{{-2, -3, 2}, {1, -1, 1}, {6, -8, 7}})
	expected := [3][3]float64{
		{-1.0 / 3.0, -5.0 / 3.0, 1.0 / 3.0},
		{1.0 / 3.0, 26.0 / 3.0, -4.0 / 3.0},
		{2.0 / 3.0, 34.0 / 3.0, -5.0 / 3.0},
	}
	inv, err := m.Inverse()
	assert.NoError(t, err)
	assert.Equal(t, expected, inv.Data(), "Inverse is incorrect")

	// singular matrix should error
	_, err = matrix.Splat(1).Inverse()
	assert.Error(t, err, "Inverting a singular matrix should error")
}

func TestIdentity(t *testing.T) {
	expected := [3][3]float64{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	assert.Equal(t, expected, matrix.Identity().Data(), "Identity is incorrect")
}

func TestNegative(t *testing.T) {
	m := matrix.New([3][3]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	expected := [3][3]float64{
		{-1, -2, -3},
		{-4, -5, -6},
		{-7, -8, -9},
	}
	assert.Equal(t, expected, m.Neg().Data(), "Neg should negate every element")
}
