package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("hello world\nos.Args %v\nArguments: %v\n", args, args[1:])
}
