package test

import (
	"math"
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

// Destructive scalar operations (Sop). nothing returned
func TestCase05(t *testing.T) {
	s := 3.0

	// 	v += s
	v, _ := vector.New(1.0, 2.0, 3.0)
	expected := []float64{v.X + s, v.Y + s, v.Z + s}
	v.Sop("+=", s)
	actual := []float64{v.X, v.Y, v.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with scalar added to eah axis")

	// 	v -= s
	expected = []float64{v.X - s, v.Y - s, v.Z - s}
	v.Sop("-=", s)
	actual = []float64{v.X, v.Y, v.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with scalar subtracting to eah axis")

	// 	v *= s
	expected = []float64{v.X * s, v.Y * s, v.Z * s}
	v.Sop("*=", s)
	actual = []float64{v.X, v.Y, v.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with scalar subtracting to eah axis")

	// 	v /= s
	expected = []float64{v.X / s, v.Y / s, v.Z / s}
	v.Sop("/=", s)
	actual = []float64{v.X, v.Y, v.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with scalar subtracting to eah axis")

	// 	v.Mag()
	expectedScalar := math.Sqrt(math.Pow(v.X, 2.0) + math.Pow(v.Y, 2.0) + math.Pow(v.Z, 2.0))
	actualScalar := v.Mag()
	assert.Equal(t, expectedScalar, actualScalar, "Magnitude of a vector should be the sqrt of the sum of the squares of each axis")

	// 	v.Normalize()
	mag := v.Mag()
	expected = []float64{v.X / mag, v.Y / mag, v.Z / mag}
	_ = v.Normalize()
	actual = []float64{v.X, v.Y, v.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with scalar subtracting to eah axis")

	// 	v.Negative()
	expected = []float64{v.X * -1.0,  v.Y * -1.0, v.Z  * -1.0}
	v.Negative()
	actual = []float64{v.X, v.Y, v.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with each axis times by -1")
}

// Destructive vector operations (Vop). nothing returned
func TestCase06(t *testing.T) {
	v1, _ := vector.New(1.0, 2.0, 3.0)
	v2, _ := vector.New(1.0, 2.0, 3.0)

	//  v1 += v2
	expected := []float64{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
	v1.Vop("+=", v2)
	actual := []float64{v1.X, v1.Y, v1.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with added vectors axis added to each axis")

	//  v1 -= v2
	expected = []float64{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
	v1.Vop("-=", v2)
	actual = []float64{v1.X, v1.Y, v1.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with substracted vectors axis substracted from each axis")

	//  v1 *= v2
	expected = []float64{v1.X * v2.X, v1.Y * v2.Y, v1.Z * v2.Z}
	v1.Vop("*=", v2)
	actual = []float64{v1.X, v1.Y, v1.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with the second vectors axis multipled by each axis")

	//  v1 /= v2
	expected = []float64{v1.X / v2.X, v1.Y / v2.Y, v1.Z / v2.Z}
	v1.Vop("/=", v2)
	actual = []float64{v1.X, v1.Y, v1.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with the second vectors axis divided by each axis")
}

// Special Object creation
func TestCase07(t *testing.T) {

	// NewX() should return 1,0,0
	expected := []float64{1.0, 0.0, 0.0}
	v := vector.NewX()
	actual := []float64{v.X, v.Y, v.Z}
	assert.Equal(t, actual, expected, "NewX() should return 1,0,0")

	// NewY() should return 0,1,0
	expected = []float64{0.0, 1.0, 0.0}
	v = vector.NewY()
	actual = []float64{v.X, v.Y, v.Z}
	assert.Equal(t, actual, expected, "NewY() should return 0,1,0")

	// NewZ() should return 0,0,1
	expected = []float64{0.0, 0.0, 1.0}
	v = vector.NewZ()
	actual = []float64{v.X, v.Y, v.Z}
	assert.Equal(t, actual, expected, "NewZ() should return 0,0,1")
}

// Vector to Vector nondestructive operations which creates a new vector
/*
With another vector (VVop)
	v1 + v2 = v3 (VpV)
	v1 - v2 = v3 (VmV)
	v1 * v2 = v3 (VtV)
	v1 / v2 = v3 (VdV)
With a scalar (VSop)
    va + s = vb (VpS)
    va - s = vb (VmS)
    va * s = vb (VtS)
    va / s = vb (VdS)
With a scalar (SVop)
	s + va = sb (SpV)
	s - va = sb (SmV)
	s * va = sb (StV)
	s / va = sb (SdV)
*/

// Other Vector operations
/*
      norm(V) = s
      unit(V) = v
	  cross(v1,vb) = vc
	  dot(va,vb) = s

*/
