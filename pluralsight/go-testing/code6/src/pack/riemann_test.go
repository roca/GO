package pack

import (
	"fmt"
)

var ri RiemannIntegrator

func ExampleRiemannIntegrator_Integrate() {

	result := ri.Integrate(0, 10, 3)

	fmt.Println(result)

	// Output:
	// 30
}
