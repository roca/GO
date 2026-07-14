package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("You did not specify a file name!")
		os.Exit(1)
	}

	fp, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, fp)

}
