package main

import (
	"fmt"

	"github.com/roca/GO/tree/staging/udemy/BuildingGoModules/toolkit"
)

func main() {
	var tools toolkit.Tools

	fmt.Println(tools.RandomString(10))
}
