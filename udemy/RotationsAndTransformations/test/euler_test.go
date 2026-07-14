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

	angles := euler.Angles{Phi: phi, Theta: theta, Si: si, Sequence: "XYZ"}
	R, _ := angles.ToDCM()

	expected := [][]float64{
		{math.Cos(theta) * math.Cos(si), math.Cos(theta) * math.Sin(si), -1.0 * math.Sin(theta)},
		{(math.Sin(phi) * math.Sin(theta) * math.Cos(si)) - (math.Cos(phi) * math.Sin(si)), (math.Sin(phi) * math.Sin(theta) * math.Sin(si)) + (math.Cos(phi) * math.Cos(si)), math.Cos(theta) * math.Sin(phi)},
		{(math.Cos(phi) * math.Sin(theta) * math.Cos(si)) + (math.Sin(phi) * math.Sin(si)), (math.Cos(phi) * math.Sin(theta) * math.Sin(si)) - (math.Sin(phi) * math.Cos(si)), math.Cos(theta) * math.Cos(phi)},
	}

	actual := R.Data()
	for i := range 3 {
		assert.InDeltaSlice(t, expected[i][:], actual[i][:], 1e-8, "XYZ DCM values incorrect")
	}
	phiActual := math.Atan2(R.M23, R.M33)
	thetaActual := -1.0 * math.Asin(R.M13)
	siActual := math.Atan2(R.M12, R.M11)
	assert.InDeltaf(t, phi, phiActual, 1e-10, "phi Values %f != %f", phi, phiActual)
	assert.InDeltaf(t, theta, thetaActual, 1e-10, "theta Values %f != %f", theta, thetaActual)
	assert.InDeltaf(t, si, siActual, 1e-10, "si Values %f != %f", si, siActual)

	xB := R.MulVec(vector.New(0.7, 1.2, -0.3))
	expectedV := []float64{0.05168617940094594, 0.6482734800319445, -1.2637523625877838}
	assert.InDeltaSlice(t, expectedV, xB.Data(), 1e-8, "R * V values incorrect")
}
