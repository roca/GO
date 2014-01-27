// bio project bio.go
package bio

import (
	"fmt"
	"io/ioutil"
	//"math"
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
	}
	head := c[0]
	tail := c[1:]
	return BytesToString(tail, acc+fmt.Sprintf("%c", head))

}

func Reverse(s string) string {
	b := []rune(s)
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func Fibonacci_old(n, k int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return Fibonacci(n-1, k) + (k * Fibonacci(n-2, k))
}

//Fibonacci Iterative
func Fibonacci(n, k int) int {
	var a, b int = 1, 1
	for i := 2; i < n; i++ {
		a, b = b, k*a+b
	}
	return b
}

func Factorial(n float64) float64 {
	loop := func(acc, n float64) float64 { return 0 }
	loop = func(acc, n float64) float64 {
		if n == 0 {
			return acc
		}
		return loop(acc*n, n-1)
	}
	return loop(1, n)
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Pascal(c, r float64) float64 {
	if c == 0 || r == 0 {
		return 1
	}
	if (r - c) == 0 {
		return 1
	}
	if r < 2 {
		return 1
	}
	return Factorial(r) / (Factorial(c) * Factorial(abs(r-c)))
}

func Fibonaccid(n, k float64) float64 {

	var a, b float64 = 1, 1

	for i := float64(2); i < n; i++ {
		if i >= k {
			a -= 1
		}
		a, b = b, a+b
	}
	return b
}

func HammingDistance(sequence1, sequence2 string) int {
	if len(sequence1) == 0 {
		return 0
	} else if sequence1[0] != sequence2[0] {
		return 1 + HammingDistance(sequence1[1:], sequence2[1:])
	}
	return HammingDistance(sequence1[1:], sequence2[1:])

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

type tRNA struct {
	Codon, Protein string
}

func codonsMap() map[string]tRNA {
	codons := make(map[string]tRNA)
	codons["UUU"] = tRNA{"UUU", "F"}
	codons["CUU"] = tRNA{"CUU", "L"}
	codons["AUU"] = tRNA{"AUU", "I"}
	codons["GUU"] = tRNA{"GUU", "V"}
	codons["UUC"] = tRNA{"UUC", "F"}
	codons["CUC"] = tRNA{"CUC", "L"}
	codons["AUC"] = tRNA{"AUC", "I"}
	codons["GUC"] = tRNA{"GUC", "V"}
	codons["UUA"] = tRNA{"UUA", "L"}
	codons["CUA"] = tRNA{"CUA", "L"}
	codons["AUA"] = tRNA{"AUA", "I"}
	codons["GUA"] = tRNA{"GUA", "V"}
	codons["UUG"] = tRNA{"UUG", "L"}
	codons["CUG"] = tRNA{"CUG", "L"}
	codons["AUG"] = tRNA{"AUG", "M"}
	codons["GUG"] = tRNA{"GUG", "V"}
	codons["UCU"] = tRNA{"UCU", "S"}
	codons["CCU"] = tRNA{"CCU", "P"}
	codons["ACU"] = tRNA{"ACU", "T"}
	codons["GCU"] = tRNA{"GCU", "A"}
	codons["UCC"] = tRNA{"UCC", "S"}
	codons["CCC"] = tRNA{"CCC", "P"}
	codons["ACC"] = tRNA{"ACC", "T"}
	codons["GCC"] = tRNA{"GCC", "A"}
	codons["UCA"] = tRNA{"UCA", "S"}
	codons["CCA"] = tRNA{"CCA", "P"}
	codons["ACA"] = tRNA{"ACA", "T"}
	codons["GCA"] = tRNA{"GCA", "A"}
	codons["UCG"] = tRNA{"UCG", "S"}
	codons["CCG"] = tRNA{"CCG", "P"}
	codons["ACG"] = tRNA{"ACG", "T"}
	codons["GCG"] = tRNA{"GCG", "A"}
	codons["UAU"] = tRNA{"UAU", "Y"}
	codons["CAU"] = tRNA{"CAU", "H"}
	codons["AAU"] = tRNA{"AAU", "N"}
	codons["GAU"] = tRNA{"GAU", "D"}
	codons["UAC"] = tRNA{"UAC", "Y"}
	codons["CAC"] = tRNA{"CAC", "H"}
	codons["AAC"] = tRNA{"AAC", "N"}
	codons["GAC"] = tRNA{"GAC", "D"}
	codons["UAA"] = tRNA{"UAA", "Stop"}
	codons["CAA"] = tRNA{"CAA", "Q"}
	codons["AAA"] = tRNA{"AAA", "K"}
	codons["GAA"] = tRNA{"GAA", "E"}
	codons["UAG"] = tRNA{"UAG", "Stop"}
	codons["CAG"] = tRNA{"CAG", "Q"}
	codons["AAG"] = tRNA{"AAG", "K"}
	codons["GAG"] = tRNA{"GAG", "E"}
	codons["UGU"] = tRNA{"UGU", "C"}
	codons["CGU"] = tRNA{"CGU", "R"}
	codons["AGU"] = tRNA{"AGU", "S"}
	codons["GGU"] = tRNA{"GGU", "G"}
	codons["UGC"] = tRNA{"UGC", "C"}
	codons["CGC"] = tRNA{"CGC", "R"}
	codons["AGC"] = tRNA{"AGC", "S"}
	codons["GGC"] = tRNA{"GGC", "G"}
	codons["UGA"] = tRNA{"UGA", "Stop"}
	codons["CGA"] = tRNA{"CGA", "R"}
	codons["AGA"] = tRNA{"AGA", "R"}
	codons["GGA"] = tRNA{"GGA", "G"}
	codons["UGG"] = tRNA{"UGG", "W"}
	codons["CGG"] = tRNA{"CGG", "R"}
	codons["AGG"] = tRNA{"AGG", "R"}
	codons["GGG"] = tRNA{"GGG", "G"}

	return codons

}
func RNAtoPROTEIN(sequence string) string {

	codons := codonsMap()

	if len(sequence) == 0 || codons[sequence[0:3]].Protein == "Stop" {
		return ""
	}
	return codons[sequence[0:3]].Protein + RNAtoPROTEIN(sequence[3:])

}

func Motifs(sequence, fragment string) []int {

	occurrences := func(sequence, fragment string, acc []int) []int { return make([]int, 1) }

	occurrences = func(sequence, fragment string, acc []int) []int {
		regexpString := fragment
		re := regexp.MustCompile(regexpString)
		found := re.FindStringIndex(sequence)

		if len(found) == 0 {
			return acc
		}

		newSequence := strings.Replace(sequence, fragment, "."+fragment[1:], 1)
		return occurrences(newSequence, fragment, append(acc, found[0]))

	}

	return occurrences(sequence, fragment, make([]int, 0))

}

func matrixColumn(matrix [][]rune, columnIndex int) []rune {
	column := make([]rune, 0)
	for _, row := range matrix {
		column = append(column, row[columnIndex])
	}
	return column
}

func matrixFromFastas(fastas []FastSequence) [][]rune {
	rows := len(fastas)
	matrix := make([][]rune, rows)
	for k, fasta := range fastas {
		matrix[k] = []rune(fasta.Sequence)
	}

	return matrix
}

func Consensus(fastaData string) string {

	fastas := FastaSequences(fastaData)
	matrix := matrixFromFastas(fastas)
	column := make([]rune, 0)
	bases := []string{"A", "C", "G", "T"}
	consensuMatrix := make([][]int, 4)
	answerText := ""

	for r, base := range bases {
		answerText = answerText + fmt.Sprintf("%s: ", base)
		consensuMatrix[r] = make([]int, len(fastas[0].Sequence))

		for k, _ := range fastas[0].Sequence {

			column = matrixColumn(matrix, k)
			consensuMatrix[r][k] = CountBaseOccurences(string(column), base)
			answerText = answerText + fmt.Sprintf("%d ", consensuMatrix[r][k])

		}
		answerText = strings.TrimSpace(answerText) + fmt.Sprintf("\n")
	}

	consensuMap := make([]string, len(fastas[0].Sequence))
	for k, _ := range fastas[0].Sequence {
		maxCount := 0
		maxBase := ""
		for r, base := range bases {
			if consensuMatrix[r][k] >= maxCount {
				maxCount = consensuMatrix[r][k]
				maxBase = base
			}
		}
		consensuMap[k] = maxBase
	}

	return fmt.Sprintf("%s\n%s", strings.Join(consensuMap, ""), answerText)
}
