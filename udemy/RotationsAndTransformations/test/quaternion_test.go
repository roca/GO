package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"udemy.com/aml/dcm"
	"udemy.com/aml/euler"
	"udemy.com/aml/quaternion"
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

	angles, _ := quat.Quat2Angles("XYZ")
	assert.InDeltaf(t, angles_xyz.Phi, angles.Phi, tolerance, "Phi %f != %f", angles_xyz.Phi, angles.Phi)
	assert.InDeltaf(t, angles_xyz.Theta, angles.Theta, tolerance, "Theta %f != %f", angles_xyz.Theta, angles.Theta)
	assert.InDeltaf(t, angles_xyz.Si, angles.Si, tolerance, "Si %f != %f", angles_xyz.Si, angles.Si)
}
