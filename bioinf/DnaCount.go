package main

import (
	//"../nf/string"
	"fmt"
	"io/ioutil"
	"os"
)

func CToGoString(c []byte) string {
	n := -1
	char := ""
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
		char := char + fmt.Sprintf("%c", n)
	}
	return char
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

	fmt.Println(CToGoString(b))
}
