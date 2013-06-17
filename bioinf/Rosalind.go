package main

import (
	"./bio"
	"fmt"
	"os"
	//"strconv"
	"strings"
)

func main() {
	args := os.Args

	data := bio.SequenceFromRosalindFile(args[1])

	sequence := strings.Replace(data, "\n", "", -1)

	got := bio.RNAtoPROTEIN(sequence)
	fmt.Printf("%s\n", got)

}
