package main

import (
	"../nf/string"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(string.Reverse(args[1]))
}
