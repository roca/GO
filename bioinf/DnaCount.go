package main

import (
	"../nf/string"
	"fmt"
	"io/ioutil"
	"os"
)

func CToGoString(c []byte) string {
	n := -1
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n+1])
}

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

	sequence := CToGoString(b[:])

	fmt.Println(string.Reverse(sequence))
}
