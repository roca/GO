package main

import "github.com/roca/GO/tree/staging/udemy/BuildingGoModules/toolkit"

func main() {
	var tools toolkit.Tools

	tools.CreateDirIfNotExist("./test-dir")
}
