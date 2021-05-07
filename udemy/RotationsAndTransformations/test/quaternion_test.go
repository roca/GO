package test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"udemy.com/aml/dcm"
	"udemy.com/aml/euler"
	"udemy.com/aml/quaternion"
	"udemy.com/aml/vector"
)

func TestAngles2Quat(t *testing.T) {
	angles_xyz := euler.New(
		dcm.DegreesToRadians(15.),
		dcm.DegreesToRadians(-35.0),
		dcm.DegreesToRadians(75.0),
	)
	tolerance := .0000000001
	expected := quaternion.Quaternion{
		S: 0.726267539555004,
		X: 0.2802526287781224,
		Y: -0.1607432936232177,
		Z: 0.6067582044000993,
	}

	quat, _ := quaternion.Angles2Quat(angles_xyz)
	assert.InDeltaf(t, expected.S, quat.S, tolerance, "quat.S %f != %f", expected.S, quat.S)
	assert.InDeltaf(t, expected.X, quat.X, tolerance, "quat.S %f != %f", expected.X, quat.X)
	assert.InDeltaf(t, expected.Y, quat.Y, tolerance, "quat.S %f != %f", expected.Y, quat.Y)
	assert.InDeltaf(t, expected.Z, quat.Z, tolerance, "quat.S %f != %f", expected.Z, quat.Z)

	angles, _ := quat.ToAngles("XYZ")
	assert.InDeltaf(t, angles_xyz.Phi, angles.Phi, tolerance, "Phi %f != %f", angles_xyz.Phi, angles.Phi)
	assert.InDeltaf(t, angles_xyz.Theta, angles.Theta, tolerance, "Theta %f != %f", angles_xyz.Theta, angles.Theta)
	assert.InDeltaf(t, angles_xyz.Si, angles.Si, tolerance, "Si %f != %f", angles_xyz.Si, angles.Si)
}

func TestExercise1(t *testing.T) {
	v, _ := vector.New(6., 12., -4.)

	phi := math.Acos(v.X / v.Mag())
	theta := math.Acos(v.Y / v.Mag())
	si := math.Acos(v.Z / v.Mag())

	r := dcm.DegreesToRadians(30)
	quat := quaternion.Quaternion{
		S: math.Cos(r / 2),
		X: math.Cos(phi) * math.Sin(r/2),
		Y: math.Cos(theta) * math.Sin(r/2),
		Z: math.Cos(si) * math.Sin(r/2),
	}

	tolerance := .001
	expected := quaternion.Quaternion{
		S: 0.9659,
		X: 0.1102,
		Y: 0.2218,
		Z: -0.0740,
	}

	assert.InDeltaf(t, expected.S, quat.S, tolerance, "quat.S %f != %f", expected.S, quat.S)
	assert.InDeltaf(t, expected.X, quat.X, tolerance, "quat.S %f != %f", expected.X, quat.X)
	assert.InDeltaf(t, expected.Y, quat.Y, tolerance, "quat.S %f != %f", expected.Y, quat.Y)
	assert.InDeltaf(t, expected.Z, quat.Z, tolerance, "quat.S %f != %f", expected.Z, quat.Z)
}
func TestSlerpInterpolate(t *testing.T) {
	angles1 := euler.New(
		dcm.DegreesToRadians(10.),
		dcm.DegreesToRadians(-20.0),
		dcm.DegreesToRadians(15.0),
	)
	angles2 := euler.New(
		dcm.DegreesToRadians(40.),
		dcm.DegreesToRadians(-60.0),
		dcm.DegreesToRadians(135.0),
	)

	quat1, _ := quaternion.Angles2Quat(angles1)
	quat2, _ := quaternion.Angles2Quat(angles2)
	quat_slerp, _ := quaternion.SlerpInterpolate(quat1, quat2, 0.5)

	tolerance := .0000000001

	angles_interp, _ := quat_slerp.ToAngles("XYZ")

	expected :=  euler.New(0.6547065473112085, -0.5673275993649903, 1.2166389498633288)

	assert.InDeltaf(t, expected.Phi, angles_interp.Phi, tolerance, "quat.S %f != %f", expected.Phi, angles_interp.Phi)
	assert.InDeltaf(t, expected.Theta, angles_interp.Theta, tolerance, "quat.S %f != %f", expected.Theta, angles_interp.Theta)
	assert.InDeltaf(t, expected.Si, angles_interp.Si, tolerance, "quat.S %f != %f", expected.Si, angles_interp.Si)

}
