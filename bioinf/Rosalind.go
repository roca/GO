package main

import (
	"./bio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args

	inputs := strings.TrimSuffix(bio.SequenceFromRosalindFile(args[1]), "\n")
	data := strings.Split(inputs, " ")

	n, err := strconv.Atoi(data[0])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	k, err := strconv.Atoi(data[1])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Printf("%d %d : %d\n", n, k, bio.Fibonacci(n-1, k))

}
