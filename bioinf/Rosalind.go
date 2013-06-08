package main

import (
	"./bio"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	sequence := bio.SequenceFromRosalindFile(args[1])

	as := bio.CountBaseOccurences(sequence, "A")
	cs := bio.CountBaseOccurences(sequence, "C")
	gs := bio.CountBaseOccurences(sequence, "G")
	ts := bio.CountBaseOccurences(sequence, "T")

	fmt.Printf("%d %d %d %d\n", as, cs, gs, ts)
}
