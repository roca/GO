package main

import (
	"fmt"

	"udemy.com/aml/dcm"
	"udemy.com/aml/euler"
)

func main() {
	attitudeAW := []float64{
		dcm.DegreesToRadians(35.0),
		dcm.DegreesToRadians(10.0),
		dcm.DegreesToRadians(75.0),
	}
	attitudeCA := []float64{
		dcm.DegreesToRadians(0.0),
		dcm.DegreesToRadians(-35.0),
		dcm.DegreesToRadians(0.0),
	}

	aw, _ := euler.New(attitudeAW[0], attitudeAW[1], attitudeAW[2], "XYZ")
	ca, _ := euler.New(attitudeCA[0], attitudeCA[1], attitudeCA[2], "XYZ")

	Raw, _ := aw.ToDCM()
	Rca, _ := ca.ToDCM()

	Rcw, _ := Rca.Mop("*", Raw)
	phi, theta, si := dcm.EulerAnglesFromRxyz(*Rcw)

	fmt.Printf("IsOrthogonal: %t\n", dcm.IsOrthogonal(*Rcw))
	fmt.Printf("Euler Angles: [%f, %f, %f] degrees\n", dcm.RadiansToDegrees(phi), dcm.RadiansToDegrees(theta), dcm.RadiansToDegrees(si))
}
