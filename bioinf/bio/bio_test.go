package bio

import (
	"fmt"
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
