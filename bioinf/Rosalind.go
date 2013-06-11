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

	for _, fasta := range fastas {
		cs := bio.CountBaseOccurences(strings.ToUpper(fasta.Sequence), "C")
		gs := bio.CountBaseOccurences(strings.ToUpper(fasta.Sequence), "G")
		gcCount := cs + gs
		fmt.Printf("%s %d\n", fasta.Header, gcCount*100.0/len(fasta.Sequence))
	}

}
