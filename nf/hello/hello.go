package main

import (
	"../string"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(string.Reverse(args[1]))
}
