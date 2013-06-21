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

	rosalindId := args[1]
	filePath := args[2]
	got := ""

	switch rosalindId {
	case "PROT":
		got = PROT(filePath)
	case "SUBS":
		got = SUBS(filePath)
	case "CONS":
		got = CONS(filePath)
	}

	fmt.Printf("%s\n", got)

}

func PROT(filePath string) string {

	data := bio.SequenceFromRosalindFile(filePath)

	sequence := strings.Replace(data, "\n", "", -1)

	return bio.RNAtoPROTEIN(sequence)

}

func SUBS(filePath string) string {

	data := strings.Split(bio.SequenceFromRosalindFile(filePath), "\n")

	sequence := strings.Join(data[:len(data)-2], "")
	fragment := data[len(data)-2]

	indices := bio.Motifs(sequence, fragment)
	got := ""
	for i, value := range indices {
		got = got + fmt.Sprintf("%d", value+1)
		if i < len(indices)-1 {
			got = got + " "
		}
	}

	return got

}

func CONS(filePath string) string {

	fastaData := bio.SequenceFromRosalindFile(filePath)
	got := bio.Consensus(fastaData)
	return got
}
