package bio

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

/*
Problem

A string is simply an ordered collection of symbols selected from some alphabet and formed into a word; the length of a string is the number of symbols that it contains.

An example of a length 21 DNA string (whose alphabet contains the symbols 'A', 'C', 'G', and 'T') is "ATGCTTCAGAAAGGTCTTACG."

Given: A DNA string s of length at most 1000 nt.

Return: Four integers (separated by spaces) counting the respective number of times that the symbols 'A', 'C', 'G', and 'T' occur in s.

Sample Dataset

AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC
Sample Output

20 12 17 21
*/

func TestCountingDNANucleotides(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC", "20 12 17 21"},
	}
	for _, c := range tests {

		sequence := c.s
		as := CountBaseOccurences(sequence, "A")
		cs := CountBaseOccurences(sequence, "C")
		gs := CountBaseOccurences(sequence, "G")
		ts := CountBaseOccurences(sequence, "T")

		got := fmt.Sprintf("%d %d %d %d", as, cs, gs, ts)
		if got != c.want {
			t.Errorf("CountBaseOccurences(%q) == %q, want %q", c.s, got, c.want)
		}

	}

}

/*

Problem

An RNA string is a string formed from the alphabet containing 'A', 'C', 'G', and 'U'.

Given a DNA string t corresponding to a coding strand, its transcribed RNA string u is formed by replacing all occurrences of 'T' in t with 'U' in u.

Given: A DNA string t having length at most 1000 nt.

Return: The transcribed RNA string of t.

Sample Dataset

GATGGAACTTGACTACGTAAATT
Sample Output

GAUGGAACUUGACUACGUAAAUU
*/
func TestTranscribingDNAintoRNA(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"GATGGAACTTGACTACGTAAATT", "GAUGGAACUUGACUACGUAAAUU"},
	}
	for _, c := range tests {

		sequence := c.s

		got := Transcribe(sequence)
		if got != c.want {
			t.Errorf("Transcribe(%q) == %q, want %q", c.s, got, c.want)
		}

	}

}

/*

Problem

In DNA strings, symbols 'A' and 'T' are complements of each other, as are 'C' and 'G'.

The reverse complement of a DNA string s is the string sc formed by reversing the symbols of s, then taking the complement of each symbol (e.g., the reverse complement of "GTCA" is "TGAC").

Given: A DNA string s of length at most 1000 bp.

Return: The reverse complement sc of s.

Sample Dataset

AAAACCCGGT
Sample Output

ACCGGGTTTT
*/
func TestComplementingaStrandofDNA(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"AAAACCCGGT", "ACCGGGTTTT"},
	}
	for _, c := range tests {

		sequence := c.s

		got := reverseComplement(sequence)
		if got != c.want {
			t.Errorf("reverseComplement(%q) == %q, want %q", c.s, got, c.want)
		}

	}

}

/*

Problem

A sequence is an ordered collection of objects (usually numbers), which are allowed to repeat. Sequences can be finite or infinite. Two examples are the finite sequence (π,−2√,0,π) and the infinite sequence of odd numbers (1,3,5,7,9,…). We use the notation an to represent the n-th term of a sequence.

A recurrence relation is a way of defining the terms of a sequence with respect to the values of previous terms. In the case of Fibonacci's rabbits from the introduction, any given month will contain the rabbits that were alive the previous month, plus any new offspring. A key observation is that the number of offspring in any month is equal to the number of rabbits that were alive two months prior. As a result, if Fn represents the number of rabbit pairs alive after the n-th month, then we obtain the Fibonacci sequence having terms Fn that are defined by the recurrence relation Fn=Fn−1+Fn−2 (with F1=F2=1 to initiate the sequence). Although the sequence bears Fibonacci's name, it was known to Indian mathematicians over two millennia ago.

When finding the n-th term of a sequence defined by a recurrence relation, we can simply use the recurrence relation to generate terms for progressively larger values of n. This problem introduces us to the computational technique of dynamic programming, which successively builds up solutions by using the answers to smaller cases.

Given: Positive integers n≤40 and k≤5.

Return: The total number of rabbit pairs that will be present after n months if each pair of reproduction-age rabbits produces a litter of k rabbit pairs in each generation (instead of only 1 pair).

Sample Dataset

5 3
Sample Output

19
*/
func TestRabbitsandRecurrenceRelations(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"5 3", "19"},
	}
	for _, c := range tests {

		inputs := strings.Split(c.s, " ")

		n, err := strconv.Atoi(inputs[0])
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}

		k, err := strconv.Atoi(inputs[1])
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}

		got := fmt.Sprintf("%d", Fibonacci(n-1, k))
		if got != c.want {
			t.Errorf("Fibonacci(%q) == %q, want %q", c.s, got, c.want)
		}

	}

}

/*

Problem

The GC-content of a DNA string is given by the percentage of symbols in the string that are 'C' or 'G'. For example, the GC-content of "AGCTATAG" is 37.5%. Note that the reverse complement of any DNA string has the same GC-content.

DNA strings must be labeled when they are consolidated into a database. A commonly used method of string labeling is called FASTA format. In this format, the string is introduced by a line that begins with '>', followed by some labeling information. Subsequent lines contain the string itself; the first line to begin with '>' indicates the label of the next string.

In Rosalind's implementation, a string in FASTA format will be labeled by the ID "Rosalind_xxxx", where "xxxx" denotes a four-digit code between 0000 and 9999.

Given: At most 10 DNA strings in FASTA format (of length at most 1 kbp each).

Return: The ID of the string having the highest GC-content, followed by the GC-content of that string. Rosalind allows for a default error of 0.001 in all decimal answers unless otherwise stated; please see the note on absolute error below.

Sample Dataset

>Rosalind_6404
CCTGCGGAAGATCGGCACTAGAATAGCCAGAACCGTTTCTCTGAGGCTTCCGGCCTTCCC
TCCCACTAATAATTCTGAGG
>Rosalind_5959
CCATCGGTAGCGCATCCTTAGTCCAATTAAGTCCCTATCCAGGCGCTCCGCCGAAGGTCT
ATATCCATTTGTCAGCAGACACGC
>Rosalind_0808
CCACCCTCGTGGTATGGCTAGGCATTCAGGAACCGGAGAACGCTTCAGACCAGCCCGGAC
TGGGAACCTGCGGGCAGTAGGTGGAAT
Sample Output

Rosalind_0808
60.919540
*/
func TestComputingGCContent(t *testing.T) {
	var tests = []struct {
		s    string
		want float64
	}{
		{">Rosalind_6404\nCCTGCGGAAGATCGGCACTAGAATAGCCAGAACCGTTTCTCTGAGGCTTCCGGCCTTCCC\nTCCCACTAATAATTCTGAGG\n>Rosalind_5959\nCCATCGGTAGCGCATCCTTAGTCCAATTAAGTCCCTATCCAGGCGCTCCGCCGAAGGTCT\nATATCCATTTGTCAGCAGACACGC\n>Rosalind_0808\nCCACCCTCGTGGTATGGCTAGGCATTCAGGAACCGGAGAACGCTTCAGACCAGCCCGGAC\nTGGGAACCTGCGGGCAGTAGGTGGAAT", 60.919540},
	}
	for _, c := range tests {

		fastas := FastaSequences(c.s)
		maxFasta := MaxGCContent(fastas)
		got := maxFasta.GCContent
		if math.Abs(got-c.want) > .001 {
			t.Errorf("MaxGCContent(%q) == %q, want %q", c.s, got, c.want)
		}

	}

}

/*
Problem


Figure 2. The Hamming distance between these two strings is 7. Mismatched symbols are colored red.
Given two strings s and t of equal length, the Hamming distance between s and t, denoted dH(s,t), is the number of corresponding symbols that differ in s and t. See Figure 2.

Given: Two DNA strings s and t of equal length (not exceeding 1 kbp).

Return: The Hamming distance dH(s,t).

Sample Dataset

GAGCCTACTAACGGGAT
CATCGTAATGACGGCCT
Sample Output

7
*/
func TestHammingDistance(t *testing.T) {
	var tests = []struct {
		s    string
		want int
	}{
		{"GAGCCTACTAACGGGAT\nCATCGTAATGACGGCCT", 7},
	}
	for _, c := range tests {

		sequence1 := strings.Split(c.s, "\n")[0]
		sequence2 := strings.Split(c.s, "\n")[1]
		got := HammingDistance(sequence1, sequence2)
		if got != c.want {
			t.Errorf("HammingDistance(%q) == %d, want %d", c.s, got, c.want)
		}

	}

}
