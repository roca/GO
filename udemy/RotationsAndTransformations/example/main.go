package main

import (
	"fmt"

	"udemy.com/aml/dcm"
)

func main() {
	phi := dcm.DegreesToRadians(-30.0)
	theta := dcm.DegreesToRadians(65.0)
	si := dcm.DegreesToRadians(-45.0)
	R, _ := dcm.Rotation(phi, theta, si)

	phiActual, thetaActual, siActual := dcm.EulerAnglesFromRotaionMatrix(R)

	d := R.Data()
	fmt.Println("Rxyz:", R.Data())
	for i := 0; i < 3; i++ {
		fmt.Println(d[i])
	}
	fmt.Printf("IsOrthogonal: %t\n", dcm.IsOrthogonal(R))
	fmt.Printf("Euler Angles: [%f, %f, %f] degrees\n", dcm.RadiansToDegrees(phiActual), dcm.RadiansToDegrees(thetaActual), dcm.RadiansToDegrees(siActual))

}
