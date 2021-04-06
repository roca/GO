package main

import (
	"fmt"

	"udemy.com/aml/dcm"
	"udemy.com/aml/euler"
)

func main() {
	phi := dcm.DegreesToRadians(-15.0)
	theta := dcm.DegreesToRadians(-105.0)
	si := dcm.DegreesToRadians(135.0)

	anglesXYZ, _ := euler.New(phi, theta, si, "XYZ")
	Rxyz, _ := anglesXYZ.ToDCM()

	phiActual, thetaActual, siActual := dcm.EulerAnglesFromRxyz(Rxyz)
	d := Rxyz.Data()
	fmt.Println("Rxyz:", Rxyz.Data())
	for i := 0; i < 3; i++ {
		fmt.Println(d[i])
	}
	fmt.Printf("IsOrthogonal: %t\n", dcm.IsOrthogonal(Rxyz))
	fmt.Printf("Euler Angles: [%f, %f, %f] degrees\n", dcm.RadiansToDegrees(phiActual), dcm.RadiansToDegrees(thetaActual), dcm.RadiansToDegrees(siActual))

	anglesZXZ, _ := euler.New(phi, theta, si, "ZXZ")
	Rzxz, _ := anglesZXZ.ToDCM()
	phiActual, thetaActual, siActual = dcm.EulerAnglesFromRzxz(Rzxz)
	d = Rzxz.Data()
	fmt.Println("Rzxz:", Rzxz.Data())
	for i := 0; i < 3; i++ {
		fmt.Println(d[i])
	}
	fmt.Printf("IsOrthogonal: %t\n", dcm.IsOrthogonal(Rzxz))
	fmt.Printf("Euler Angles: [%f, %f, %f] degrees\n", dcm.RadiansToDegrees(phiActual), dcm.RadiansToDegrees(thetaActual), dcm.RadiansToDegrees(siActual))

	phiActual, thetaActual, siActual = dcm.EulerAnglesFromRzxz(Rxyz)
	fmt.Printf("Attitude for ZXZ: [%f, %f, %f] degrees\n", dcm.RadiansToDegrees(phiActual), dcm.RadiansToDegrees(thetaActual), dcm.RadiansToDegrees(siActual))

	anglesZXZActual, _ := euler.New(phiActual, thetaActual, siActual, "ZXZ")
	Rzxz, _ = anglesZXZActual.ToDCM()
	phiActual, thetaActual, siActual = dcm.EulerAnglesFromRxyz(Rzxz)
	fmt.Printf("Attitude for ZXZ: [%f, %f, %f] degrees\n", dcm.RadiansToDegrees(phiActual), dcm.RadiansToDegrees(thetaActual), dcm.RadiansToDegrees(siActual))

}
