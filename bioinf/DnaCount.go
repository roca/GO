package main

import (
	"./bio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args

	// read whole the file
	b, err := ioutil.ReadFile(args[1])
	if err != nil {
		panic(err)
	}

	// write whole the body
	err = ioutil.WriteFile("output.txt", b, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println(bio.BytesToString(b, ""))
}
