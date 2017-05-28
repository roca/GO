package main

import (
	"fmt"

	"github.com/GOCODE/pluralsight/go-testing/code4/02_Check/src/pack"
)

func main() {

	pi := pack.PolyIntegrator{}

	fmt.Println(pi.Integrate(0, 10, 3))
	fmt.Println(pi.Integrate(0, 10, 1, 0))

	ri := pack.RiemannIntegrator{}

	fmt.Println(ri.Integrate(0, 10, 3))
	fmt.Println(ri.Integrate(0, 10, 1, 0))
}
