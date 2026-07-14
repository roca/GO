package test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"udemy.com/aml/dcm"
	"udemy.com/aml/matrix"
	"udemy.com/aml/vector"
)

func TestRotationX(t *testing.T) {
	theta := 30. * math.Pi / 180.
	m := dcm.RotationX(theta)
	expected := [3][3]float64{
		{1, 0, 0},
		{0, math.Cos(theta), math.Sin(theta)},
		{0, -math.Sin(theta), math.Cos(theta)},
	}
	actual := m.Data()
	for i := range 3 {
		assert.InDeltaSlice(t, expected[i][:], actual[i][:], 1e-8, "RotationX Matrix values incorrect")
	}
	xB := m.MulVec(vector.New(0.7, 1.2, -0.3))
	expectedV := []float64{0.7, 0.88923048, -.85980762}
	assert.InDeltaSlice(t, expectedV, xB.Data(), 1e-8, "RotationX * V values incorrect")
}

func TestRotationY(t *testing.T) {
	theta := 30. * math.Pi / 180.
	m := dcm.RotationY(theta)
	expected := [3][3]float64{
		{math.Cos(theta), 0, -math.Sin(theta)},
		{0, 1, 0},
		{math.Sin(theta), 0, math.Cos(theta)},
	}
	actual := m.Data()
	for i := range 3 {
		assert.InDeltaSlice(t, expected[i][:], actual[i][:], 1e-8, "RotationY Matrix values incorrect")
	}
	xB := m.MulVec(vector.New(0.7, 1.2, -0.3))
	expectedV := []float64{0.756217782, 1.2, 0.09019237886}
	assert.InDeltaSlice(t, expectedV, xB.Data(), 1e-8, "RotationY * V values incorrect")
}

func TestRotationZ(t *testing.T) {
	theta := 30. * math.Pi / 180.
	m := dcm.RotationZ(theta)
	expected := [3][3]float64{
		{math.Cos(theta), math.Sin(theta), 0},
		{-math.Sin(theta), math.Cos(theta), 0},
		{0, 0, 1},
	}
	actual := m.Data()
	for i := range 3 {
		assert.InDeltaSlice(t, expected[i][:], actual[i][:], 1e-8, "RotationZ Matrix values incorrect")
	}
	xB := m.MulVec(vector.New(0.7, 1.2, -0.3))
	expectedV := []float64{1.2062177826, 0.689230484541, -0.3}
	assert.InDeltaSlice(t, expectedV, xB.Data(), 1e-8, "RotationZ * V values incorrect")
}

func TestIsOrthogonal(t *testing.T) {
	m := matrix.New([3][3]float64{
		{1.0 / 3.0, -2.0 / 3.0, 2.0 / 3.0},
		{2.0 / 3.0, -1.0 / 3.0, -2.0 / 3.0},
		{2.0 / 3.0, 2.0 / 3.0, 1.0 / 3.0},
	})
	assert.True(t, dcm.IsOrthogonal(m), "This Matrix should be orthogonal")
}

func TestIsNotOrthogonal(t *testing.T) {
	m := matrix.New([3][3]float64{
		{1, 2, 2},
		{2, 1, 2},
		{2, 2, 1},
	})
	assert.False(t, dcm.IsOrthogonal(m), "This Matrix should not be orthogonal")
}

func TestNormalize(t *testing.T) {
	m := matrix.New([3][3]float64{
		{1.0 / 3.0, -2.0 / 3.0, 2.0 / 3.0},
		{2.0 / 3.0, -1.0 / 3.0, -2.0 / 3.0},
		{2.0 / 3.0, 2.0 / 3.0, 1.0 / 3.0},
	})
	dcm.Normalize(&m)
	assert.True(t, dcm.IsOrthogonal(m), "This Matrix should be orthogonal")
}

func TestKinematicRatesFromBodyRates(t *testing.T) {
	R := matrix.Identity()
	rates := vector.New(1, 0, 0)
	for range 100 {
		RDot := dcm.KinematicRatesFromBodyRates(R, rates)
		R = dcm.Integrate(R, RDot, .01)
	}
	assert.True(t, dcm.IsOrthogonal(R, 1e-6), "Integrated DCM should remain orthogonal")
}
