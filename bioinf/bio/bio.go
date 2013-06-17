// bio project bio.go
package bio

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func CountBaseOccurences(sequence, base string) int {
	return strings.Count(strings.ToUpper(sequence), strings.ToUpper(base))
}

func Transcribe(sequence string) string {

	transcriptionMap := map[string]string{"t": "u", "T": "U"}

	transcribedSequence := sequence
	for key, value := range transcriptionMap {
		transcribedSequence = strings.Replace(transcribedSequence, key, value, -1)
	}
	return transcribedSequence

}

func ReverseComplement(sequence string) string {

	complementMap := map[string]string{"A": "T", "C": "G", "T": "A", "G": "C"}

	reverseComplemented := ""
	bs := []rune(sequence)
	for i := 0; i < len(bs); i++ {
		reverseComplemented = complementMap[string(bs[i])] + reverseComplemented

	}
	return reverseComplemented

}

func regSplit(text string, delimeter string) []string {
	reg := regexp.MustCompile(delimeter)
	indexes := reg.FindAllStringIndex(text, -1)
	laststart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = text[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = text[laststart:len(text)]
	return result
}

func FastaHeaders(text string) []string {
	reg := regexp.MustCompile(">.*\n")
	return reg.FindAllString(text, -1)
}

type FastSequence struct {
	Header, Sequence string
	GCContent        float64
}

func FastaSequences(text string) []FastSequence {
	headers := FastaHeaders(text)
	sequences := regSplit(text, ">.*\n")

	fastas := make([]FastSequence, len(headers), cap(headers))
	for i, header := range headers {
		fastas[i].Header = strings.Replace(header, ">", "", -1)
		fastas[i].Sequence = strings.Replace(sequences[i+1], "\n", "", -1)
		cs := float64(CountBaseOccurences(strings.ToUpper(fastas[i].Sequence), "C"))
		gs := float64(CountBaseOccurences(strings.ToUpper(fastas[i].Sequence), "G"))
		fastas[i].GCContent = ((cs + gs) / float64(len(fastas[i].Sequence))) * 100.0

	}

	return fastas
}
func MaxGCContent(fastas []FastSequence) FastSequence {
	var maxFasta FastSequence
	for _, fasta := range fastas {
		if fasta.GCContent > maxFasta.GCContent {
			maxFasta = fasta
		}
	}
	return maxFasta
}

func SequenceFromRosalindFile(filePath string) string {

	// read whole the file
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// write whole the body
	err = ioutil.WriteFile("output.txt", b, 0644)
	if err != nil {
		panic(err)
	}

	return BytesToString(b, "")

}

func BytesToString(c []byte, acc string) string {

	if len(c) == 0 {
		return acc
	} else {
		head := c[0]
		tail := c[1:]
		return BytesToString(tail, acc+fmt.Sprintf("%c", head))
	}
}

func Reverse(s string) string {
	b := []rune(s)
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func Fibonacci(n int, k int) int {
	fib := 0
	if n == 1 || n == 0 {
		fib = 1
	} else {
		fib = Fibonacci(n-1, k) + (k * Fibonacci(n-2, k))
	}
	return fib
}

func HammingDistance(sequence1, sequence2 string) int {
	if len(sequence1) == 0 {
		return 0
	} else if sequence1[0] != sequence2[0] {
		return 1 + HammingDistance(sequence1[1:], sequence2[1:])
	} else {
		return HammingDistance(sequence1[1:], sequence2[1:])
	}

}

func ChanceOfDominantPhenotype(dominant int, mixed int, recessive int) float64 {
	domnt := float64(dominant)
	mix := float64(mixed)
	recssv := float64(recessive)
	total := domnt + mix + recssv
	side1 := domnt / total
	side2 := (mix / total) * (.5) * (1.0 + (domnt / (total - 1.0)) + (.5)*((mix-1.0)/(total-1.0)))
	side3 := (recssv/total)*(domnt/(total-1.0)) + (recssv/total)*(mix/(total-1.0))*(.5)
	return (side1 + side2 + side3)
}

/*
UUU F      CUU L      AUU I      GUU V
UUC F      CUC L      AUC I      GUC V
UUA L      CUA L      AUA I      GUA V
UUG L      CUG L      AUG M      GUG V
UCU S      CCU P      ACU T      GCU A
UCC S      CCC P      ACC T      GCC A
UCA S      CCA P      ACA T      GCA A
UCG S      CCG P      ACG T      GCG A
UAU Y      CAU H      AAU N      GAU D
UAC Y      CAC H      AAC N      GAC D
UAA Stop   CAA Q      AAA K      GAA E
UAG Stop   CAG Q      AAG K      GAG E
UGU C      CGU R      AGU S      GGU G
UGC C      CGC R      AGC S      GGC G
UGA Stop   CGA R      AGA R      GGA G
UGG W      CGG R      AGG R      GGG G
*/

type Codon struct {
	Dna, Protein string
}

func codonsMap() map[string]Codon {
	codons := make(map[string]Codon)
	codons["UUU"] = Codon{"UUU", "F"}
	codons["CUU"] = Codon{"CUU", "L"}
	codons["AUU"] = Codon{"AUU", "I"}
	codons["GUU"] = Codon{"GUU", "V"}
	codons["UUC"] = Codon{"UUC", "F"}
	codons["CUC"] = Codon{"CUC", "L"}
	codons["AUC"] = Codon{"AUC", "I"}
	codons["GUC"] = Codon{"GUC", "V"}
	codons["UUA"] = Codon{"UUA", "L"}
	codons["CUA"] = Codon{"CUA", "L"}
	codons["AUA"] = Codon{"AUA", "I"}
	codons["GUA"] = Codon{"GUA", "V"}
	codons["UUG"] = Codon{"UUG", "L"}
	codons["CUG"] = Codon{"CUG", "L"}
	codons["AUG"] = Codon{"AUG", "M"}
	codons["GUG"] = Codon{"GUG", "V"}
	codons["UCU"] = Codon{"UCU", "S"}
	codons["CCU"] = Codon{"CCU", "P"}
	codons["ACU"] = Codon{"ACU", "T"}
	codons["GCU"] = Codon{"GCU", "A"}
	codons["UCC"] = Codon{"UCC", "S"}
	codons["CCC"] = Codon{"CCC", "P"}
	codons["ACC"] = Codon{"ACC", "T"}
	codons["GCC"] = Codon{"GCC", "A"}
	codons["UCA"] = Codon{"UCA", "S"}
	codons["CCA"] = Codon{"CCA", "P"}
	codons["ACA"] = Codon{"ACA", "T"}
	codons["GCA"] = Codon{"GCA", "A"}
	codons["UCG"] = Codon{"UCG", "S"}
	codons["CCG"] = Codon{"CCG", "P"}
	codons["ACG"] = Codon{"ACG", "T"}
	codons["GCG"] = Codon{"GCG", "A"}
	codons["UAU"] = Codon{"UAU", "Y"}
	codons["CAU"] = Codon{"CAU", "H"}
	codons["AAU"] = Codon{"AAU", "N"}
	codons["GAU"] = Codon{"GAU", "D"}
	codons["UAC"] = Codon{"UAC", "Y"}
	codons["CAC"] = Codon{"CAC", "H"}
	codons["AAC"] = Codon{"AAC", "N"}
	codons["GAC"] = Codon{"GAC", "D"}
	codons["UAA"] = Codon{"UAA", "Stop"}
	codons["CAA"] = Codon{"CAA", "Q"}
	codons["AAA"] = Codon{"AAA", "K"}
	codons["GAA"] = Codon{"GAA", "E"}
	codons["UAG"] = Codon{"UAG", "Stop"}
	codons["CAG"] = Codon{"CAG", "Q"}
	codons["AAG"] = Codon{"AAG", "K"}
	codons["GAG"] = Codon{"GAG", "E"}
	codons["UGU"] = Codon{"UGU", "C"}
	codons["CGU"] = Codon{"CGU", "R"}
	codons["AGU"] = Codon{"AGU", "S"}
	codons["GGU"] = Codon{"GGU", "G"}
	codons["UGC"] = Codon{"UGC", "C"}
	codons["CGC"] = Codon{"CGC", "R"}
	codons["AGC"] = Codon{"AGC", "S"}
	codons["GGC"] = Codon{"GGC", "G"}
	codons["UGA"] = Codon{"UGA", "Stop"}
	codons["CGA"] = Codon{"CGA", "R"}
	codons["AGA"] = Codon{"AGA", "R"}
	codons["GGA"] = Codon{"GGA", "G"}
	codons["UGG"] = Codon{"UGG", "W"}
	codons["CGG"] = Codon{"CGG", "R"}
	codons["AGG"] = Codon{"AGG", "R"}
	codons["GGG"] = Codon{"GGG", "G"}

	return codons

}
func RNAtoPROTEIN(sequence string) string {

	codons := codonsMap()

	if len(sequence) == 0 || codons[sequence[0:3]].Protein == "Stop" {
		return ""
	} else {
		return codons[sequence[0:3]].Protein + RNAtoPROTEIN(sequence[3:])
	}

}
