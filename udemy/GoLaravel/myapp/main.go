package main

import (
	"fmt"

	"github.com/roca/celeritas"
)

func main() {
	result := celeritas.TestFunc(1, 2)
	fmt.Println(result)

	result  = celeritas.TestFunc2(1, 2)
	fmt.Println(result)

	result = celeritas.TestFunc3(1, 2)
	fmt.Println(result)
}
