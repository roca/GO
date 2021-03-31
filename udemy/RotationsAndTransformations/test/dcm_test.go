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
func TestXYZRotation(t *testing.T) {
	phi := dcm.DegreesToRadians(-45.0)
	theta := dcm.DegreesToRadians(-75.0)
	si := dcm.DegreesToRadians(78.0)
	R, _ := dcm.XYZRotation(phi, theta, si)

	expected := [][]float64{
		{math.Cos(theta) * math.Cos(si), math.Cos(theta) * math.Sin(si), -1.0 * math.Sin(theta)},
		{(math.Sin(phi) * math.Sin(theta) * math.Cos(si)) - (math.Cos(phi) * math.Sin(si)), (math.Sin(phi) * math.Sin(theta) * math.Sin(si)) + (math.Cos(phi) * math.Cos(si)), math.Cos(theta) * math.Sin(phi)},
		{(math.Cos(phi) * math.Sin(theta) * math.Cos(si)) + (math.Sin(phi) * math.Sin(si)), (math.Cos(phi) * math.Sin(theta) * math.Sin(si)) - (math.Sin(phi) * math.Cos(si)), math.Cos(theta) * math.Cos(phi)},
	}

	actual := R.Data()
	for i := 0; i < 3; i++ {
		b := assert.InDeltaSlice(t, expected[i], actual[i], .00000001)
		assert.Equal(t, true, b, "RotationZ Matrix values incorrect")
	}
	phiActual := math.Atan2(R.M23, R.M33)
	thetaActual := -1.0 * math.Asin(R.M13)
	siActual := math.Atan2(R.M12, R.M11)
	assert.InDeltaf(t, phi, phiActual, .0000000001, "phi Values %f != %f", phi, phiActual)
	assert.InDeltaf(t, theta, thetaActual, .0000000001, "theta Values %f != %f", theta, thetaActual)
	assert.InDeltaf(t, si, siActual, .0000000001, "si Values %f != %f", si, siActual)

	x_a, _ := vector.New(0.7, 1.2, -0.3)
	x_b, _ := R.Vop("*", x_a)

	expectedV := []float64{0.05168617940094594, 0.6482734800319445, -1.2637523625877838}
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
		v, _ := vector.New([]float64{1., 0.0, 0.0})
		RDot, _ := dcm.KinematicRatesFromBodyRates(&R, &v)
		R, _ = dcm.Intergrate(&R, &RDot, .01)
		fmt.Println(R)
	}

}
