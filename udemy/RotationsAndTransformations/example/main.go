package main

import (
	"fmt"
	"math"

	"udemy.com/aml/dcm"
)

func main() {
	phi := -45. * math.Pi / 180.
	theta := -75. * math.Pi / 180.
	si := 78.0 * math.Pi / 180.
	m, _ := dcm.Rotation(phi, theta, si)

	phiActual := math.Atan(m.M23 / m.M33)
	thetaActual := -1.0 * math.Asin(m.M13)
	siActual := math.Atan(m.M12 / m.M11)

	d := m.Data()
	fmt.Println("Rxyz:", m.Data())
	for i := 0; i<3;i++ {
		fmt.Println(d[i])
	}
	fmt.Printf("IsOrthogonal: %t\n",dcm.IsOrthogonal(m))
	fmt.Printf("Euler Angles: [%f, %f, %f]\n", radiansToDegrees(phiActual), radiansToDegrees(thetaActual), radiansToDegrees(siActual))

}
func radiansToDegrees(radian float64) (degrees float64) {
	degrees = radian * 180.0 / math.Pi
	return
}
