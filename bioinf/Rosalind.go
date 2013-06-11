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
		cs := float64(bio.CountBaseOccurences(strings.ToUpper(fasta.Sequence), "C"))
		gs := float64(bio.CountBaseOccurences(strings.ToUpper(fasta.Sequence), "G"))
		gcCount := cs + gs
		length := len(fasta.Sequence)
		fmt.Printf("%s %4.4g %d %4.10g\n", fasta, gcCount, length, (gcCount/float64(length))*100.0)
	}

}
