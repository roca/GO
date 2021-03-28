package test

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"udemy.com/aml/dcm"
	"udemy.com/aml/matrix"
	"udemy.com/aml/vector"
)

func TestRotationX(t *testing.T) {
	theta := 30. * math.Pi / 180.
	m, _ := dcm.RotationX(theta)
	expected := [][]float64{
		{1., 0., 0.},
		{0., math.Cos(theta), math.Sin(theta)},
		{0., -1.0 * math.Sin(theta), math.Cos(theta)},
	}
	actual := m.Data()
	for i := 0; i < 3; i++ {
		b := assert.InDeltaSlice(t, expected[i], actual[i], .00000001)
		assert.Equal(t, true, b, "RotationX Matrix values incorrect")
	}
	x_a, _ := vector.New(0.7, 1.2, -0.3)
	x_b, _ := m.Vop("*", x_a)

	expectedV := []float64{0.7, 0.88923048, -.85980762}
	actualV := x_b.Data()
	for i := 0; i < 3; i++ {
		b := assert.InDelta(t, expectedV[i], actualV[i], .00000001)
		assert.Equal(t, true, b, "RotationX * V values incorrect")
	}

}
func TestRotationY(t *testing.T) {
	theta := 30. * math.Pi / 180.
	m, _ := dcm.RotationY(theta)
	expected := [][]float64{
		{math.Cos(theta), 0., -1.0 * math.Sin(theta)},
		{0., 1.0, 0.},
		{math.Sin(theta), 0., math.Cos(theta)},
	}
	actual := m.Data()
	for i := 0; i < 3; i++ {
		b := assert.InDeltaSlice(t, expected[i], actual[i], .00000001)
		assert.Equal(t, true, b, "RotationY Matrix values incorrect")
	}
	x_a, _ := vector.New(0.7, 1.2, -0.3)
	x_b, _ := m.Vop("*", x_a)

	expectedV := []float64{0.756217782, 1.2, 0.09019237886}
	actualV := x_b.Data()
	for i := 0; i < 3; i++ {
		b := assert.InDelta(t, expectedV[i], actualV[i], .00000001)
		assert.Equal(t, true, b, "RotationX * V values incorrect")
	}

}

func TestRotationZ(t *testing.T) {
	theta := 30. * math.Pi / 180.
	m, _ := dcm.RotationZ(theta)
	expected := [][]float64{
		{math.Cos(theta), math.Sin(theta), 0.},
		{-1.0 * math.Sin(theta), math.Cos(theta), 0.},
		{0., 0., 1.},
	}
	actual := m.Data()
	for i := 0; i < 3; i++ {
		b := assert.InDeltaSlice(t, expected[i], actual[i], .00000001)
		assert.Equal(t, true, b, "RotationZ Matrix values incorrect")
	}
	x_a, _ := vector.New(0.7, 1.2, -0.3)
	x_b, _ := m.Vop("*", x_a)

	expectedV := []float64{1.2062177826, 0.689230484541, -0.3}
	actualV := x_b.Data()
	for i := 0; i < 3; i++ {
		b := assert.InDelta(t, expectedV[i], actualV[i], .00000001)
		assert.Equal(t, true, b, "RotationX * V values incorrect")
	}

}
func TestRotation(t *testing.T) {
	theta := 30. * math.Pi / 180.
	m, _ := dcm.Rotation(0., 0., theta)
	expected := [][]float64{
		{math.Cos(theta), math.Sin(theta), 0.},
		{-1.0 * math.Sin(theta), math.Cos(theta), 0.},
		{0., 0., 1.},
	}
	actual := m.Data()
	for i := 0; i < 3; i++ {
		b := assert.InDeltaSlice(t, expected[i], actual[i], .00000001)
		assert.Equal(t, true, b, "RotationZ Matrix values incorrect")
	}
	x_a, _ := vector.New(0.7, 1.2, -0.3)
	x_b, _ := m.Vop("*", x_a)

	expectedV := []float64{1.2062177826, 0.689230484541, -0.3}
	actualV := x_b.Data()
	for i := 0; i < 3; i++ {
		b := assert.InDelta(t, expectedV[i], actualV[i], .00000001)
		assert.Equal(t, true, b, "RotationX * V values incorrect")
	}

}
func TestIsOrthogonal(t *testing.T) {
	m, _ := matrix.New([][]float64{
		{1.0 / 3.0, -2.0 / 3.0, 2.0 / 3.0},
		{2.0 / 3.0, -1.0 / 3.0, -2.0 / 3.0},
		{2.0 / 3.0, 2.0 / 3.0, 1.0 / 3.0},
	})
	assert.True(t, dcm.IsOrthogonal(m), "This Matrix should be orthogonal")
}
func TestIsNotOrthogonal(t *testing.T) {
	m, _ := matrix.New([][]float64{
		{1.0, 2.0, 2.0},
		{2.0, 1.0, 2.0},
		{2.0, 2.0, 1.0},
	})
	assert.False(t, dcm.IsOrthogonal(m), "This Matrix should not be orthogonal")
}
func TestNormaliz(t *testing.T) {
	m, _ := matrix.New([][]float64{
		{1.0 / 3.0, -2.0 / 3.0, 2.0 / 3.0},
		{2.0 / 3.0, -1.0 / 3.0, -2.0 / 3.0},
		{2.0 / 3.0, 2.0 / 3.0, 1.0 / 3.0},
	})
	_ = dcm.Normalize(&m)
	assert.True(t, dcm.IsOrthogonal(m), "This Matrix should be orthogonal")
}

func TestKinematicRatesFromBodyRates(t *testing.T) {
	R := matrix.Identity()
	for i := 0; i < 100; i++ {
		v, _ := vector.New([]float64{1.,0.0,0.0})
		RDot,_ := dcm.KinematicRatesFromBodyRates(&R, &v)
		R,_ = dcm.Intergrate(&R, &RDot, .01)
		fmt.Println(R)
	}


}
