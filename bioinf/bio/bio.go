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

func reverseComplement(sequence string) string {

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

/*
	def chanceOfDominantPhenotype(dominant:Int, mixed:Int, recessive:Int):BigDecimal = {
		val total = Double.box(dominant + mixed + recessive)
		val side1 = dominant/total
		val side2 = (mixed/total)*(1/2d)*(1 + (dominant/(total-1)) + (1/2d)*((mixed-1)/(total-1)))
		val side3 = (recessive/total)*(dominant/(total-1)) + (recessive/total)*(mixed/(total-1))*(1/2d)
		BigDecimal.apply(side1 + side2 + side3).setScale(5, RoundingMode.HALF_UP)
	}
*/
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
