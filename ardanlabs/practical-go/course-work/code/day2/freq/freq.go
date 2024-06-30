package main

// Q: What is the most common word (ignoring case) in sherlock.txt?

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

// Who's on first -> [Who s on first]
var wordRe *regexp.Regexp = regexp.MustCompile(`\w+`)

//var wordRe *regexp.Regexp = regexp.MustCompile(`[a-zA-Z]+`)

/* Will run before main
func init() {

}
*/

func main() {
	// read the file
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	freqs, err := wordFrequency(file)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	for word, count := range freqs {
		fmt.Printf("%s: %d\n", word, count)
	}

	mostFreqWord, mostFreqCount, err := maxWord(freqs)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("Most common word: %s (%d)\n", mostFreqWord, mostFreqCount)
	// count the frequency of each word
	// print the most common word
}

func sortByFrequenceCount(freqs map[string]int) []string {
	words := make([]string, 0, len(freqs))
	for word := range freqs {
		words = append(words, word)
	}
	sort.Slice(words, func(i, j int) bool {
		return freqs[words[i]] > freqs[words[j]]
	})
	return words
}

func maxWord(freqs map[string]int) (string, int, error) {
	if len(freqs) == 0 {
		return "", 0, fmt.Errorf("empty map")
	}

	maxN, maxW := 0, ""
	for word, count := range freqs {
		if count > maxN {
			maxN, maxW = count, word
		}
	}
	return maxW, maxN, nil
}

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	freqs := make(map[string]int)
	for s.Scan() {
		words := wordRe.FindAllString(s.Text(), -1) // current line
		for _, w := range words {
			freqs[strings.ToLower(w)]++
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	return freqs, nil
}
