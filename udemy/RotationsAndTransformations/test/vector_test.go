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

/*
	Destructive scalar operations (Sop). nothing return
	(v *Vector) Sop(operation string, value float64) (*Vector, error)
*/
func TestCase05(t *testing.T) {
	s := 3.0

	// 	v += s
	v, _ := vector.New(1.0, 2.0, 3.0)
	expected := []float64{v.X + s, v.Y + s, v.Z + s}
	v.Sop("+=", s)
	actual := []float64{v.X, v.Y, v.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with scalar added to each axis")

	// 	v + s
	expected = []float64{v.X + s, v.Y + s, v.Z + s}
	u, _ := v.Sop("+", s)
	actual = []float64{u.X, u.Y, u.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with scalar added to each axis")

	// 	v -= s
	expected = []float64{v.X - s, v.Y - s, v.Z - s}
	v.Sop("-=", s)
	actual = []float64{v.X, v.Y, v.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with scalar subtracted from each axis")

	// 	v - s
	expected = []float64{v.X - s, v.Y - s, v.Z - s}
	w, _ := v.Sop("-", s)
	actual = []float64{w.X, w.Y, w.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with scalar subtracted from each axis")

	// 	v *= s
	expected = []float64{v.X * s, v.Y * s, v.Z * s}
	v.Sop("*=", s)
	actual = []float64{v.X, v.Y, v.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with scalar multiplied by each axis")

	// 	v * s
	expected = []float64{v.X * s, v.Y * s, v.Z * s}
	x, _ := v.Sop("*", s)
	actual = []float64{x.X, x.Y, x.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with scalar multiplied by each axis")

	// 	v /= s
	expected = []float64{v.X / s, v.Y / s, v.Z / s}
	v.Sop("/=", s)
	actual = []float64{v.X, v.Y, v.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with each axis divided by scalar")

	// 	v / s
	expected = []float64{v.X / s, v.Y / s, v.Z / s}
	y, _ := v.Sop("/", s)
	actual = []float64{y.X, y.Y, y.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with each axis divided by scalar")

	// 	v ? s
	expected = []float64{0.0, 0.0, 0.0}
	z, e := v.Sop("?", s)
	actual = []float64{z.X, z.Y, z.Z}
	assert.Equal(t, actual, expected, "Vector should a nil object for its typ")
	assert.NotNil(t, e, "This unknown operations should raise error")

	// 	v.Mag()
	expectedScalar := math.Sqrt(math.Pow(v.X, 2.0) + math.Pow(v.Y, 2.0) + math.Pow(v.Z, 2.0))
	actualScalar := v.Mag()
	assert.Equal(t, expectedScalar, actualScalar, "Magnitude of a vector should be the sqrt of the sum of the squares of each axis")

	// 	v.Normalize()
	mag := v.Mag()
	expected = []float64{v.X / mag, v.Y / mag, v.Z / mag}
	v.Normalize()
	actual = []float64{v.X, v.Y, v.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with scalar subtracting to each axis")

	// 	v.Negative()
	expected = []float64{v.X * -1.0, v.Y * -1.0, v.Z * -1.0}
	v.Negative()
	actual = []float64{v.X, v.Y, v.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with each axis times by -1")
}

/*
 	Destructive vector operations (Vop). nothing returned
	(v *Vector) Vop(operation string, u Vector) (*Vector, error)
*/
func TestCase06(t *testing.T) {
	v1, _ := vector.New(1.0, 2.0, 3.0)
	v2, _ := vector.New(1.0, 2.0, 3.0)

	//  v1 += v2
	expected := []float64{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
	v1.Vop("+=", v2)
	actual := []float64{v1.X, v1.Y, v1.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with added vectors axis added to each axis")

	//  v1 + v2
	expected = []float64{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
	v1.Vop("+", v2)
	actual = []float64{v1.X, v1.Y, v1.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with added vectors axis added to each axis")

	//  v1 -= v2
	expected = []float64{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
	v1.Vop("-=", v2)
	actual = []float64{v1.X, v1.Y, v1.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with substracted vectors axis substracted from each axis")

	//  v1 - v2
	expected = []float64{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
	v1.Vop("-", v2)
	actual = []float64{v1.X, v1.Y, v1.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with substracted vectors axis substracted from each axis")

	//  v1 *= v2
	expected = []float64{v1.X * v2.X, v1.Y * v2.Y, v1.Z * v2.Z}
	v1.Vop("*=", v2)
	actual = []float64{v1.X, v1.Y, v1.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with the second vectors axis multipled by each axis")

	//  v1 * v2
	expected = []float64{v1.X * v2.X, v1.Y * v2.Y, v1.Z * v2.Z}
	v1.Vop("*", v2)
	actual = []float64{v1.X, v1.Y, v1.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with the second vectors axis multipled by each axis")

	//  v1 /= v2
	expected = []float64{v1.X / v2.X, v1.Y / v2.Y, v1.Z / v2.Z}
	v1.Vop("/=", v2)
	actual = []float64{v1.X, v1.Y, v1.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with the second vectors axis divided by each axis")

	//  v1 / v2
	expected = []float64{v1.X / v2.X, v1.Y / v2.Y, v1.Z / v2.Z}
	v1.Vop("/", v2)
	actual = []float64{v1.X, v1.Y, v1.Z}
	assert.Equal(t, actual, expected, "Vector should be altered with the second vectors axis divided by each axis")

	// 	v ? s
	expected = []float64{0.0, 0.0, 0.0}
	z, e := v1.Vop("?", v2)
	actual = []float64{z.X, z.Y, z.Z}
	assert.Equal(t, actual, expected, "Vector should a nil object for its type")
	assert.NotNil(t, e, "This unknown operations should raise error")
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
// Cross
func TestCase08(t *testing.T) {
	v1, _ := vector.New(2.0, -5.0, 4.0)
	v2, _ := vector.New(6.0, 2.0, -8.0)
	expected := []float64{32.0, 40.0, 34.0}
	v := vector.Cross(v1, v2)
	actual := []float64{v.X, v.Y, v.Z}
	assert.Equal(t, actual, expected, "Cross(v1,v2) is inccorrect")
}

// Dot
func TestCase09(t *testing.T) {
	v1, _ := vector.New(2.0, -5.0, 4.0)
	v2, _ := vector.New(6.0, 2.0, -8.0)
	expected := -30.0
	actual := vector.Dot(v1, v2)
	assert.Equal(t, actual, expected, "Dot(v1,v2) is inccorrect")
}

// Unit
func TestCase10(t *testing.T) {
	v, _ := vector.New(2.0, -5.0, 4.0)
	expected := []float64{v.X / v.Norm(), v.Y / v.Norm(), v.Z / v.Norm()}
	u := vector.Unit(v)
	actual := []float64{u.X, u.Y, u.Z}
	assert.Equal(t, actual, expected, "Unit(v) is inccorrect")
}
