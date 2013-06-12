package main

import (
	"./bio"
	"fmt"
	"os"
	"strings"
	//"strconv"
)

func main() {
	args := os.Args

	data := bio.SequenceFromRosalindFile(args[1])

	sequence1 := strings.Split(data, "\n")[0]
	sequence2 := strings.Split(data, "\n")[1]
	got := bio.HammingDistance(sequence1, sequence2)
	fmt.Printf("%d\n", got)

}
