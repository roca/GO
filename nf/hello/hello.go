package main

import (
	"fmt"
	"github.com/nf/string"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(string.Reverse(args[1]))
}
