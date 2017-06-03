package pack

import "fmt"

func ExamplePolyIntegrator_Integrate() {
	pi := PolyIntegrator{}

	result := pi.Integrate(0, 10, 3)

	fmt.Println(result)

	// Output:
	// 30
}
func ExamplePolyIntegrator_Integrate_line() {
	pi := PolyIntegrator{}

	result := pi.Integrate(0, 10, 1, 0)

	fmt.Println(result)

	// Output:
	// 50
}
