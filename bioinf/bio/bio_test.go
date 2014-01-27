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
Problem DNA

A string is simply an ordered collection of symbols selected from some alphabet and formed into a word; the length of a string is the number of symbols that it contains.

An example of a length 21 DNA string (whose alphabet contains the symbols 'A', 'C', 'G', and 'T') is "ATGCTTCAGAAAGGTCTTACG."

Given: A DNA string s of length at most 1000 nt.

Return: Four integers (separated by spaces) counting the respective number of times that the symbols 'A', 'C', 'G', and 'T' occur in s.

Sample Dataset

AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC
Sample Output

20 12 17 21
*/

func TestDNA(t *testing.T) {
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

Problem RNA

An RNA string is a string formed from the alphabet containing 'A', 'C', 'G', and 'U'.

Given a DNA string t corresponding to a coding strand, its transcribed RNA string u is formed by replacing all occurrences of 'T' in t with 'U' in u.

Given: A DNA string t having length at most 1000 nt.

Return: The transcribed RNA string of t.

Sample Dataset

GATGGAACTTGACTACGTAAATT
Sample Output

GAUGGAACUUGACUACGUAAAUU
*/
func TestRNA(t *testing.T) {
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

Problem REVC

In DNA strings, symbols 'A' and 'T' are complements of each other, as are 'C' and 'G'.

The reverse complement of a DNA string s is the string sc formed by reversing the symbols of s, then taking the complement of each symbol (e.g., the reverse complement of "GTCA" is "TGAC").

Given: A DNA string s of length at most 1000 bp.

Return: The reverse complement sc of s.

Sample Dataset

AAAACCCGGT
Sample Output

ACCGGGTTTT
*/
func TestREVC(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"AAAACCCGGT", "ACCGGGTTTT"},
	}
	for _, c := range tests {

		sequence := c.s

		got := ReverseComplement(sequence)
		if got != c.want {
			t.Errorf("ReverseComplement(%q) == %q, want %q", c.s, got, c.want)
		}

	}

}

/*

Problem FIB

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
func TestFIB(t *testing.T) {
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

		got := fmt.Sprintf("%d", Fibonacci(n, k))
		if got != c.want {
			t.Errorf("Fibonacci(%q) == %q, want %q", c.s, got, c.want)
		}

	}

}

/*

Problem GC

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
func TestGC(t *testing.T) {
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
Problem HAMM


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
func TestHAMM(t *testing.T) {
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

/*

Problem IPRB


Figure 2. The probability of any outcome (leaf) in a probability tree diagram is given by the product of probabilities from the start of the tree to the outcome. For example, the probability that X is blue and Y is blue is equal to (2/5)(1/4), or 1/10.
Probability is the mathematical study of randomly occurring phenomena. We will model such a phenomenon with a random variable, which is simply a variable that can take a number of different distinct outcomes depending on the result of an underlying random process.

For example, say that we have a bag containing 3 red balls and 2 blue balls. If we let X represent the random variable corresponding to the color of a drawn ball, then the probability of each of the two outcomes is given by Pr(X=red)=35 and Pr(X=blue)=25.

Random variables can be combined to yield new random variables. Returning to the ball example, let Y model the color of a second ball drawn from the bag (without replacing the first ball). The probability of Y being red depends on whether the first ball was red or blue. To represent all outcomes of X and Y, we therefore use a probability tree diagram. This branching diagram represents all possible individual probabilities for X and Y, with outcomes at the endpoints ("leaves") of the tree. The probability of any outcome is given by the product of probabilities along the path from the beginning of the tree; see Figure 2 for an illustrative example.

An event is simply a collection of outcomes. Because outcomes are distinct, the probability of an event can be written as the sum of the probabilities of its constituent outcomes. For our colored ball example, let A be the event "Y is blue." Pr(A) is equal to the sum of the probabilities of two different outcomes: Pr(X=blue and Y=blue)+Pr(X=red and Y=blue), or 310+110=25 (see Figure 2 above).

Given: Three positive integers k, m, and n, representing a population containing k+m+n organisms: k individuals are homozygous dominant for a factor, m are heterozygous, and n are homozygous recessive.

Return: The probability that two randomly selected mating organisms will produce an individual possessing a dominant allele (and thus displaying the dominant phenotype). Assume that any two organisms can mate.

Sample Dataset

2 2 2
Sample Output

0.78333

*/
func TestIPRB(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"2 2 2", "0.78333"},
	}
	for _, c := range tests {

		inputs := strings.Split(c.s, " ")

		k, err := strconv.Atoi(inputs[0])
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}

		m, err := strconv.Atoi(inputs[1])
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}

		n, err := strconv.Atoi(inputs[2])
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}

		got := fmt.Sprintf("%1.5g", ChanceOfDominantPhenotype(k, m, n))
		if got != c.want {
			t.Errorf("ChanceOfDominantPhenotype(%q) == %q, want %q", c.s, got, c.want)
		}

	}

}

/*
Problem PROT

The 20 commonly occurring amino acids are abbreviated by using 20 letters from the English alphabet (all letters except for B, J, O, U, X, and Z). Protein strings are constructed from these 20 symbols. Henceforth, the term genetic string will incorporate protein strings along with DNA strings and RNA strings.

The RNA codon table dictates the details regarding the encoding of specific codons into the amino acid alphabet.

Given: An RNA string s corresponding to a strand of mRNA (of length at most 10 kbp).

Return: The protein string encoded by s.

Sample Dataset

AUGGCCAUGGCGCCCAGAACUGAGAUCAAUAGUACCCGUAUUAACGGGUGA
Sample Output

MAMAPRTEINSTRING
*/

func TestPROT(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"AUGGCCAUGGCGCCCAGAACUGAGAUCAAUAGUACCCGUAUUAACGGGUGA", "MAMAPRTEINSTRING"},
	}
	for _, c := range tests {

		sequence := c.s

		got := RNAtoPROTEIN(sequence)
		if got != c.want {
			t.Errorf("RNAtoPROTEIN(%q) == %q, want %q", c.s, got, c.want)
		}

	}

}

/*
Problem SUBS

Given two strings s and t, t is a substring of s if t is contained as a contiguous collection of symbols in s (as a result, t must be no longer than s).

The position of a symbol in a string is the total number of symbols found to its left, including itself (e.g., the positions of all occurrences of 'U' in "AUGCUUCAGAAAGGUCUUACG" are 2, 5, 6, 15, 17, and 18). The symbol at position i of s is denoted by s[i].

A substring of s can be represented as s[j:k], where j and k represent the starting and ending positions of the substring in s; for example, if s = "AUGCUUCAGAAAGGUCUUACG", then s[2:5] = "UGCU".

The location of a substring s[j:k] is its beginning position j; note that t will have multiple locations in s if it occurs more than once as a substring of s (see the Sample below).

Given: Two DNA strings s and t (each of length at most 1 kbp).

Return: All locations of t as a substring of s.

Sample Dataset

GATATATGCATATACTT
ATAT
Sample Output

2 4 10
*/

func TestSUBS(t *testing.T) {
	var tests = []struct {
		s1, s2, want string
	}{
		{"GATATATGCATATACTT", "ATAT", "2 4 10"},
	}
	for _, c := range tests {

		sequence := c.s1
		fragment := c.s2

		indices := Motifs(sequence, fragment)
		got := ""
		for i, value := range indices {
			got = got + fmt.Sprintf("%d", value+1)
			if i < len(indices)-1 {
				got = got + " "
			}
		}

		if got != c.want {
			t.Errorf("Motifs(%q,%q) == %q, want %q", c.s1, c.s2, got, c.want)
		}

	}

}

/*

Problem CONS

A matrix is a rectangular table of values divided into rows and columns. An m×n matrix has m rows and n columns. Given a matrix A, we write Ai,j to indicate the value found at the intersection of row i and column j.

Say that we have a collection of DNA strings, all having the same length n. Their profile matrix is a 4×n matrix P in which P1,j represents the number of times that 'A' occurs in the jth position of one of the strings, P2,j represents the number of times that C occurs in the jth position, and so on (see below).

A consensus string c is a string of length n formed from our collection by taking the most common symbol at each position; the jth symbol of c therefore corresponds to the symbol having the maximum value in the j-th column of the profile matrix. Of course, there may be more than one most common symbol, leading to multiple possible consensus strings.

A T C C A G C T
G G G C A A C T
A T G G A T C T
DNA Strings	A A G C A A C C
T T G G A A C T
A T G C C A T T
A T G G C A C T
A   5 1 0 0 5 5 0 0
Profile	C   0 0 1 4 2 0 6 1
G   1 1 6 3 0 1 0 0
T   1 5 0 0 0 1 1 6
Consensus	A T G C A A C T
Given: A collection of at most 10 DNA strings of equal length (at most 1 kbp) in FASTA format.

Return: A consensus string and profile matrix for the collection. (If several possible consensus strings exist, then you may return any one of them.)

Sample Dataset

>Rosalind_1
ATCCAGCT
>Rosalind_2
GGGCAACT
>Rosalind_3
ATGGATCT
>Rosalind_4
AAGCAACC
>Rosalind_5
TTGGAACT
>Rosalind_6
ATGCCATT
>Rosalind_7
ATGGCACT
Sample Output

ATGCAACT
A: 5 1 0 0 5 5 0 0
C: 0 0 1 4 2 0 6 1
G: 1 1 6 3 0 1 0 0
T: 1 5 0 0 0 1 1 6
*/
func TestCONS(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{">Rosalind_1\nATCCAGCT\n>Rosalind_2\nGGGCAACT\n>Rosalind_3\nATGGATCT\n>Rosalind_4\nAAGCAACC\n>Rosalind_5\nTTGGAACT\n>Rosalind_6\nATGCCATT\n>Rosalind_7\nATGGCACT",
			"ATGCAACT\nA: 5 1 0 0 5 5 0 0\nC: 0 0 1 4 2 0 6 1\nG: 1 1 6 3 0 1 0 0\nT: 1 5 0 0 0 1 1 6\n"},
	}
	for _, c := range tests {

		fastaData := c.s

		got := Consensus(fastaData)

		if got != c.want {
			t.Errorf("Consensus(%q) ==\n        %q,\nwant %q", c.s, got, c.want)
		}

	}

}

/*
Problem FIBD


Figure 4. A figure illustrating the propagation of Fibonacci's rabbits if they die after three months.
Recall the definition of the Fibonacci numbers from “Rabbits and Recurrence Relations”, which followed the recurrence relation Fn=Fn−1+Fn−2 and assumed that each pair of rabbits reaches maturity in one month and produces a single pair of offspring (one male, one female) each subsequent month.

Our aim is to somehow modify this recurrence relation to achieve a dynamic programming solution in the case that all rabbits die out after a fixed number of months. See Figure 4 for a depiction of a rabbit tree in which rabbits live for three months (meaning that they reproduce only twice before dying).

Given: Positive integers n≤100 and m≤20.

Return: The total number of pairs of rabbits that will remain after the n-th month if all rabbits live for m months.

Sample Dataset

6 3
Sample Output

4
*/
func TestFIBD(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"6 3", "4"},
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

		got := fmt.Sprintf("%g", Fibonaccid(float64(n), float64(k)))
		if got != c.want {
			t.Errorf("Fibonacci(%q) == %2.0q, want %2.0q", c.s, got, c.want)
		}

	}

}
