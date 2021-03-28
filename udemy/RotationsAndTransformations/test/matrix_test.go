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
	v1, _ := vector.New(1., 2., 3.)
	v2, _ := vector.New(4., 5., 6.)
	v3, _ := vector.New(7., 8., 9.)
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
	v1, _ := vector.New(1., 2., 3.)
	v2, _ := vector.New(4., 5., 6.)
	v3, _ := vector.New(7., 8., 9.)
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
func TestCopy(t *testing.T) {
	m, _ := matrix.New([][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}, {7.0, 8.0, 9.0}})
	u, _ := m.Copy()
	u.M11 = 0.0

	if &u == &m {
		t.Errorf("Addresses should not be equal %g %g", u, &m)
	}
	if u.M11 == m.M11 {
		t.Errorf("Values should not be equal %g %g", u.M11, m.M11)
	}
	assert.Equal(t, 1.0, m.M11, "Incorrect value")
	assert.Equal(t, 0.0, u.M11, "Incorrect value")

}
func TestCopyPointer(t *testing.T) {
	m, _ := matrix.New([][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}, {7.0, 8.0, 9.0}})
	u, _ := m.CopyPointer()
	u.M11 = 0.0

	if u == &m {
		t.Errorf("Addresses should not be equal %g %g", u, &m)
	}
	if u.M11 == m.M11 {
		t.Errorf("Values should not be equal %g %g", u.M11, m.M11)
	}
	assert.Equal(t, 1.0, m.M11, "Incorrect value")
	assert.Equal(t, 0.0, u.M11, "Incorrect value")

}

//Operations with a matrix
func TestAdditionWithMatrix(t *testing.T) {
	m1, _ := matrix.New([][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}, {7.0, 8.0, 9.0}})
	m2, _ := matrix.New([][]float64{{1.5, 2.5, 3.5}, {4.5, 5.5, 6.5}, {7.5, 8.5, 9.5}})
	_, _ = m1.Mop("+=", m2)
	expected := [][]float64{
		{2.5, 4.5, 6.5},
		{8.5, 10.5, 12.5},
		{14.5, 16.5, 18.5},
	}
	actual := m1.Data()
	for i := 0; i < 3; i++ {
		b := assert.InDeltaSlice(t, expected[i], actual[i], .000000000000001)
		assert.Equal(t, true, b, "Matrix addition values incorrect")
	}
}

func TestAdditionWithMatrixNonDestructive(t *testing.T) {
	m1, _ := matrix.New([][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}, {7.0, 8.0, 9.0}})
	m2, _ := matrix.New([][]float64{{1.5, 2.5, 3.5}, {4.5, 5.5, 6.5}, {7.5, 8.5, 9.5}})
	m3, _ := m1.Mop("+", m2)
	expected := [][]float64{
		{2.5, 4.5, 6.5},
		{8.5, 10.5, 12.5},
		{14.5, 16.5, 18.5},
	}
	actual := m3.Data()
	for i := 0; i < 3; i++ {
		b := assert.InDeltaSlice(t, expected[i], actual[i], .000000000000001)
		assert.Equal(t, true, b, "Matrix addition values incorrect")
	}
	if m3 == &m1 {
		t.Errorf("Addresses should not be equal %g %g", m3, &m1)
	}
	if m3.M11 == m1.M11 {
		t.Errorf("Values should not be equal %g %g", m3.M11, m1.M11)
	}
	assert.Equal(t, 2.5, m3.M11, "Incorrect value")
	assert.Equal(t, 1.0, m1.M11, "Incorrect value")

}
func TestSubtractionWithMatrix(t *testing.T) {
	m1, _ := matrix.New([][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}, {7.0, 8.0, 9.0}})
	m2, _ := matrix.New([][]float64{{1.5, 2.5, 3.5}, {4.5, 5.5, 6.5}, {7.5, 8.5, 9.5}})
	_, _ = m1.Mop("-=", m2)
	expected := [][]float64{
		{-0.5, -0.5, -0.5},
		{-0.5, -0.5, -0.5},
		{-0.5, -0.5, -0.5},
	}
	actual := m1.Data()
	for i := 0; i < 3; i++ {
		b := assert.InDeltaSlice(t, expected[i], actual[i], .000000000000001)
		assert.Equal(t, true, b, "Matrix subtraction values incorrect")
	}
}
func TestMultiplicationWithMatrix(t *testing.T) {
	m1, _ := matrix.New([][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}, {7.0, 8.0, 9.0}})
	m2, _ := matrix.New([][]float64{{1.5, 2.5, 3.5}, {4.5, 5.5, 6.5}, {7.5, 8.5, 9.5}})
	_, _ = m1.Mop("*=", m2)
	expected := [][]float64{
		{33.0, 39.0, 45.0},
		{73.5, 88.5, 103.5},
		{114.0, 138.0, 162.0},
	}
	actual := m1.Data()
	for i := 0; i < 3; i++ {
		b := assert.InDeltaSlice(t, expected[i], actual[i], .000000000000001)
		assert.Equal(t, true, b, "Matrix multiplication values incorrect")
	}
}
func TestDivisionWithMatrix(t *testing.T) {
	m1, _ := matrix.New([][]float64{{-2.0, -3.0, 2.0}, {1.0, 0.0, 1.0}, {6.0, -8.0, 7.0}})
	m2, _ := matrix.New([][]float64{{-2.0, 2.0, 3.0}, {-1.0, 1.0, 3.0}, {2.0, 0.0, -1.0}})
	_, _ = m1.Mop("/=", m2)
	expected := [][]float64{
		{-17.0 / 6.0, 8. / 3., -5.0 / 2.0},
		{-0.5, 1.0, 0.5},
		{-10.0, 12.0, -1.0},
	}
	actual := m1.Data()
	for i := 0; i < 3; i++ {
		b := assert.InDeltaSlice(t, expected[i], actual[i], .000000000000001)
		assert.Equal(t, true, b, "Matrix division values incorrect")
	}
}

//Operations with a vector
func TestMultiplicationWithVector(t *testing.T) {
	m, _ := matrix.New([][]float64{
		{-2.0, -3.0, 2.0},
		{1.0, 0.0, 1.0},
		{6.0, -8.0, 7.0},
	})
	v, _ := vector.New([]float64{-2.0, 2.0, 3.0})
	v2, _ := m.Vop("*", v)
	expected := []float64{4.0, 1.0, -7.0}
	actual := []float64{v2.X, v2.Y, v2.Z}
	assert.Equal(t, expected, actual, "M * V operation is incorrect")
}

/*
	Destructive scalar operations (Sop). nothing return
	(m *Matrix) Sop(operation string, value float64) (*Matrix, error)
*/
func TestAdditionWithScalar(t *testing.T) {
	m, _ := matrix.New([][]float64{
		{-2.0, -3.0, 2.0},
		{1.0, -1.0, 1.0},
		{6.0, -8.0, 7.0},
	})
	u, _ := m.Sop("+=", 0.5)
	actual := u.Data()
	expected := [][]float64{
		{-2.0 + 0.5, -3.0 + 0.5, 2.0 + 0.5},
		{1.0 + 0.5, -1.0 + 0.5, 1.0 + 0.5},
		{6.0 + 0.5, -8.0 + 0.5, 7.0 + 0.5},
	}
	assert.Equal(t, expected, actual, "Matrix += S calculations is inncorrect")
}
func TestSubtractionWithScalar(t *testing.T) {
	m, _ := matrix.New([][]float64{
		{-2.0, -3.0, 2.0},
		{1.0, -1.0, 1.0},
		{6.0, -8.0, 7.0},
	})
	u, _ := m.Sop("-=", 0.5)
	actual := u.Data()
	expected := [][]float64{
		{-2.0 - 0.5, -3.0 - 0.5, 2.0 - 0.5},
		{1.0 - 0.5, -1.0 - 0.5, 1.0 - 0.5},
		{6.0 - 0.5, -8.0 - 0.5, 7.0 - 0.5},
	}
	assert.Equal(t, expected, actual, "Matrix -= S calculations is inncorrect")
}
func TestMultiplicationWithScalar(t *testing.T) {
	m, _ := matrix.New([][]float64{
		{-2.0, -3.0, 2.0},
		{1.0, -1.0, 1.0},
		{6.0, -8.0, 7.0},
	})
	u, _ := m.Sop("*=", 0.5)
	actual := u.Data()
	expected := [][]float64{
		{-2.0 * 0.5, -3.0 * 0.5, 2.0 * 0.5},
		{1.0 * 0.5, -1.0 * 0.5, 1.0 * 0.5},
		{6.0 * 0.5, -8.0 * 0.5, 7.0 * 0.5},
	}
	assert.Equal(t, expected, actual, "Matrix *= S calculations is inncorrect")
}
func TestDivisionWithScalar(t *testing.T) {
	m, _ := matrix.New([][]float64{
		{-2.0, -3.0, 2.0},
		{1.0, -1.0, 1.0},
		{6.0, -8.0, 7.0},
	})
	u, _ := m.Sop("/=", 0.5)
	actual := u.Data()
	expected := [][]float64{
		{-2.0 / 0.5, -3.0 / 0.5, 2.0 / 0.5},
		{1.0 / 0.5, -1.0 / 0.5, 1.0 / 0.5},
		{6.0 / 0.5, -8.0 / 0.5, 7.0 / 0.5},
	}
	assert.Equal(t, expected, actual, "Matrix /= S calculations is inncorrect")
}

//Special operations
func TestDiag(t *testing.T) {
	// With a Vector
	v, _ := vector.New([]float64{-2.0, -3.0, 2.0})
	m, _ := matrix.Diag(v)
	expected := [][]float64{
		{-2., 0., 0.},
		{0., -3., 0.},
		{0., 0., 2.},
	}
	actual := m.Data()
	assert.Equal(t, expected, actual, "Diagonal valuse should be the original vector")

	u, _ := matrix.New([][]float64{
		{-2.0, -3.0, 2.0},
		{1.0, -1.0, 1.0},
		{6.0, -8.0, 7.0},
	})
	v, _ = matrix.DiagV(u)
	expected2 := []float64{-2., -1., 7.}
	actual2 := []float64{v.X, v.Y, v.Z}
	assert.Equal(t, expected2, actual2, "Diagonal valuse should be the original vector")
}
func TestTranspose(t *testing.T) {
	m, _ := matrix.New([][]float64{
		{-2.0, -3.0, 2.0},
		{1.0, -1.0, 1.0},
		{6.0, -8.0, 7.0},
	})
	u, _ := matrix.Transpose(m)
	expected := [][]float64{
		{-2.0, 1.0, 6.0},
		{-3.0, -1.0, -8.0},
		{2.0, 1.0, 7.0},
	}
	actual := u.Data()
	assert.Equal(t, actual, expected, "Transpose of this matrix is incorrect")
}
func TestDeterminant(t *testing.T) {
	m, _ := matrix.New([][]float64{
		{-2.0, -3.0, 2.0},
		{1.0, -1.0, 1.0},
		{6.0, -8.0, 7.0},
	})
	expected := -3.0
	actual, _ := matrix.Determinant(m)
	assert.Equal(t, expected, actual, "Determinant should be %f", expected)
}
func TestInverse(t *testing.T) {
	m, _ := matrix.New([][]float64{
		{-2.0, -3.0, 2.0},
		{1.0, -1.0, 1.0},
		{6.0, -8.0, 7.0},
	})
	expected := [][]float64{
		{-1. / 3., -5. / 3., 1. / 3.},
		{1. / 3., 26. / 3., -4. / 3.},
		{2. / 3., 34. / 3., -5. / 3.},
	}
	inverseM, _ := m.Inverse()
	actual := inverseM.Data()

	assert.Equal(t, actual, expected, "Inverse is incorrect")
}
func TestIdentity(t *testing.T) {

	expected := [][]float64{
		{1., 0., 0.},
		{0., 1., 0.},
		{0., 0., 1.},
	}
	identityM := matrix.Identity()
	actual := identityM.Data()

	assert.Equal(t, actual, expected, "All valuse should along the diaganal should be equal to 1.0. All others 0.0")
}
func TestNegative(t *testing.T) {
	m, _ := matrix.New([][]float64{
		{1., 2., 3.},
		{4., 5., 6.},
		{7., 8., 9.},
	})
	expected := [][]float64{
		{-1., -2., -3.},
		{-4., -5., -6.},
		{-7., -8., -9.},
	}
	negativeM, _ := m.Negate()
	actual := negativeM.Data()

	assert.Equal(t, actual, expected, "All valuse should be the negative of the original")
}
func TestEpsilon(t *testing.T) {
	assert.Equal(t, matrix.Epsilon(1, 2, 3), 1, "e(1,1,1) should equal 1.0")
	assert.Equal(t, matrix.Epsilon(2, 3, 1), 1, "e(1,1,1) should equal 1.0")
	assert.Equal(t, matrix.Epsilon(3, 1, 2), 1, "e(1,1,1) should equal 1.0")

	assert.Equal(t, matrix.Epsilon(3, 2, 1), -1, "e(3,2,1) should equal -1.0")
	assert.Equal(t, matrix.Epsilon(1, 3, 2), -1, "e(1,3,2) should equal -1.0")
	assert.Equal(t, matrix.Epsilon(2, 1, 3), -1, "e(2,1,3) should equal -1.0")

}
