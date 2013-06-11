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

	var fastas []bio.FastSequence = bio.FastaSequences(data)

	for i := 0; i < len(fastas); i++ {
		cs := bio.CountBaseOccurences(strings.ToUpper(fastas[i].Sequence), "C")
		gs := bio.CountBaseOccurences(strings.ToUpper(fastas[i].Sequence), "G")
		gcCount := cs + gs
		fmt.Printf("%s %d\n", fastas[i], gcCount*100.0/len(fastas[i].Sequence))
	}

}
