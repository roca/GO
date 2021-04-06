package test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"udemy.com/aml/dcm"
	"udemy.com/aml/euler"
	"udemy.com/aml/vector"
)

func TestToXYZDCM(t *testing.T) {
	phi := dcm.DegreesToRadians(-45.0)
	theta := dcm.DegreesToRadians(-75.0)
	si := dcm.DegreesToRadians(78.0)

	angles, _ := euler.New(phi, theta, si, "XYZ")
	R, _ := angles.ToDCM()

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
