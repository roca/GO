package test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"udemy.com/aml/vector"
)

// Case 01: Default values for Vector should be 0.0
func TestCase01(t *testing.T) {
	v := vector.Vector{}

	expected := []float64{0.0, 0.0, 0.0}
	actual := []float64{v.X, v.Y, v.Z}

	assert.Equal(t, expected, actual, "Default values for Vector should be 0.0 float64")
}

// Case 02: Splat should initialize every component with a scalar value
func TestCase02(t *testing.T) {
	v := vector.Splat(3.0)

	expected := []float64{3.0, 3.0, 3.0}
	actual := []float64{v.X, v.Y, v.Z}

	assert.Equal(t, expected, actual, "Splat should initialize with scalar value")
}

// Case 03: New should initialize with three values
func TestCase03(t *testing.T) {
	v := vector.New(1.0, 2.0, 3.0)

	expected := []float64{1.0, 2.0, 3.0}
	actual := []float64{v.X, v.Y, v.Z}

	assert.Equal(t, expected, actual, "New should initialize with three values")
}

// Case 04: FromSlice should initialize with a slice of three values
func TestCase04(t *testing.T) {
	v, err := vector.FromSlice([]float64{3.0, 2.0, 1.0})
	assert.NoError(t, err)

	expected := []float64{3.0, 2.0, 1.0}
	actual := []float64{v.X, v.Y, v.Z}

	assert.Equal(t, expected, actual, "FromSlice should initialize with slice of three values")

	// Wrong length should error
	_, err = vector.FromSlice([]float64{1.0, 2.0})
	assert.Error(t, err, "FromSlice with wrong length should error")
}

// Case 05: Scalar operations return new vectors (value semantics)
func TestCase05(t *testing.T) {
	s := 3.0
	v := vector.New(1.0, 2.0, 3.0)

	// Scale
	scaled := v.Scale(s)
	assert.Equal(t, []float64{v.X * s, v.Y * s, v.Z * s}, []float64{scaled.X, scaled.Y, scaled.Z}, "Scale multiplies each axis")

	// receiver is unchanged (value semantics)
	assert.Equal(t, []float64{1.0, 2.0, 3.0}, []float64{v.X, v.Y, v.Z}, "Scale must not mutate the receiver")

	// Neg
	neg := v.Neg()
	assert.Equal(t, []float64{-v.X, -v.Y, -v.Z}, []float64{neg.X, neg.Y, neg.Z}, "Neg negates each axis")

	// Mag
	expectedMag := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	assert.Equal(t, expectedMag, v.Mag(), "Mag is sqrt of the sum of squares")

	// Normalize
	unit, err := v.Normalize()
	assert.NoError(t, err)
	mag := v.Mag()
	assert.Equal(t, []float64{v.X / mag, v.Y / mag, v.Z / mag}, []float64{unit.X, unit.Y, unit.Z}, "Normalize divides each axis by the magnitude")

	// Normalize of zero vector errors
	_, err = vector.Vector{}.Normalize()
	assert.Error(t, err, "Normalizing a zero-magnitude vector should error")
}

// Case 06: Vector-to-vector operations return new vectors
func TestCase06(t *testing.T) {
	v1 := vector.New(1.0, 2.0, 3.0)
	v2 := vector.New(4.0, 5.0, 6.0)

	add := v1.Add(v2)
	assert.Equal(t, []float64{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}, []float64{add.X, add.Y, add.Z}, "Add sums each axis")

	sub := v1.Sub(v2)
	assert.Equal(t, []float64{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}, []float64{sub.X, sub.Y, sub.Z}, "Sub subtracts each axis")

	mul := v1.Mul(v2)
	assert.Equal(t, []float64{v1.X * v2.X, v1.Y * v2.Y, v1.Z * v2.Z}, []float64{mul.X, mul.Y, mul.Z}, "Mul multiplies each axis")

	div := v1.Div(v2)
	assert.Equal(t, []float64{v1.X / v2.X, v1.Y / v2.Y, v1.Z / v2.Z}, []float64{div.X, div.Y, div.Z}, "Div divides each axis")

	// receiver unchanged
	assert.Equal(t, []float64{1.0, 2.0, 3.0}, []float64{v1.X, v1.Y, v1.Z}, "operations must not mutate the receiver")
}

// Case 07: Axis unit vectors
func TestCase07(t *testing.T) {
	x := vector.UnitX()
	assert.Equal(t, []float64{1.0, 0.0, 0.0}, []float64{x.X, x.Y, x.Z}, "UnitX() should return 1,0,0")

	y := vector.UnitY()
	assert.Equal(t, []float64{0.0, 1.0, 0.0}, []float64{y.X, y.Y, y.Z}, "UnitY() should return 0,1,0")

	z := vector.UnitZ()
	assert.Equal(t, []float64{0.0, 0.0, 1.0}, []float64{z.X, z.Y, z.Z}, "UnitZ() should return 0,0,1")
}

// Case 08: Cross product
func TestCase08(t *testing.T) {
	v1 := vector.New(2.0, -5.0, 4.0)
	v2 := vector.New(6.0, 2.0, -8.0)
	expected := []float64{32.0, 40.0, 34.0}
	v := vector.Cross(v1, v2)
	assert.Equal(t, expected, []float64{v.X, v.Y, v.Z}, "Cross(v1,v2) is incorrect")
}

// Case 09: Dot product
func TestCase09(t *testing.T) {
	v1 := vector.New(2.0, -5.0, 4.0)
	v2 := vector.New(6.0, 2.0, -8.0)
	expected := -30.0
	assert.Equal(t, expected, vector.Dot(v1, v2), "Dot(v1,v2) is incorrect")
}

// Case 10: Unit vector
func TestCase10(t *testing.T) {
	v := vector.New(2.0, -5.0, 4.0)
	mag := v.Norm()
	expected := []float64{v.X / mag, v.Y / mag, v.Z / mag}
	u := vector.Unit(v)
	assert.Equal(t, expected, []float64{u.X, u.Y, u.Z}, "Unit(v) is incorrect")
}
