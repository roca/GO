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

	data := bio.SequenceFromRosalindFile(args[1])

	inputs := strings.Split(strings.Replace(data, "\n", "", -1), " ")

	k, err := strconv.Atoi(inputs[0])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	m, err := strconv.Atoi(inputs[1])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	n, err := strconv.Atoi(inputs[2])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	got := bio.ChanceOfDominantPhenotype(k, m, n)
	fmt.Printf("%1.4g\n", got)

}
