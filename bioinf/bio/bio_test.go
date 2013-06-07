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
