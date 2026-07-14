package main

import (
	"fmt"

	"udemy.com/aml/euler"
)

func main() {

	angles := euler.New(0.1, -0.3, -0.5)
	dcm, _ := angles.ToDCM()

	recoveredAngles, _ := euler.DcmToAngles(dcm, angles.Sequence)

	fmt.Printf("Angles: %v\n", angles)
	fmt.Printf("Recovered Angles: %v", recoveredAngles)
}
