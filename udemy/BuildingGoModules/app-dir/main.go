package main

import "github.com/roca/go-toolkit"

func main() {
	var tools toolkit.Tools

	tools.CreateDirIfNotExist("./test-dir")
}
