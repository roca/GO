package main

import (
	"./bio"
	"fmt"
	"os"
	//"strconv"
)

func main() {
	args := os.Args

	data := bio.SequenceFromRosalindFile(args[1])

	var fastas []bio.FastSequence = bio.FastaSequences(data)
	var maxFasta = bio.MaxGCContent(fastas)
	fmt.Printf("%s%4.6g\n", maxFasta.Header, maxFasta.GCContent)

}
