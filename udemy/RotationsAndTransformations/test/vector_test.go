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
}

// 	v.normalize

// Vector destructive operations (Vop). nothing returned
//  v1 += v2
//  v1 -= v2
// 	v1 *= v2
// 	v1 /= v2

// Special Object creation
//    NewX(v ...interface) should return 1,0,0
//    NewY(v ...interface) should return 0,1,0
//    NewZ(v ...interface) should return 0,0,1

// Vector to Vector nondestructive operations which creates a new vector
/*
With another vector
	v1 + v2 = v3 (VpV)
	v1 - v2 = v3 (VmV)
	v1 * v2 = v3 (VtV)
	v1 / v2 = v3 (VdV)
With a scalar
    va + s = vb (VpS)
    va - s = vb (VmS)
    va * s = vb (VtS)
    va / s = vb (VdS)
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
