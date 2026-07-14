package main

import (
	"bufio"
	"fmt"
	"log"
	"maps"
	"os"
	"regexp"
	"slices"
	"sort"
	"strings"
)

// What are the N most common words in sherlock.txt

var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

func main() {
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wordFrequencies := make(map[string]int)

	for scanner.Scan() {
		words := wordRe.FindAllString(scanner.Text(), -1)

		for _, word := range words {
			wordFrequencies[strings.ToLower(word)]++
		}
	}
	// This method give these top 5: [{the 5816} {and 3089} {i 3038} {to 2825} {of 2780}]

	// scanner.Split(bufio.ScanWords)

	// // Set the scanner to split by words
	// scanner.Split(bufio.ScanWords)

	// for scanner.Scan() {
	// 	word := scanner.Text()
	// 	wordFrequencies[strings.ToLower(word)]++
	// 	//fmt.Println(word)
	// }
	// This method give these top 5: [{the 5703} {and 2882} {of 2758} {to 2720} {a 2648}]

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %v", err)
	}

	top := topN(wordFrequencies, 10)

	fmt.Println(top)
}

func topN(freq map[string]int, n int) []string {
	words := slices.Collect(maps.Keys(freq))
	sort.Slice(words, func(i, j int) bool {
		wi, wj := words[i], words[j]
		// Sort in reverse order
		return freq[wi] > freq[wj]
	})

	n = min(n, len(words))
	return words[:n]
}

func topN2(freq map[string]int, n int) []string {
	pairList := rankByWordCount(freq)
	words := []string{}

	for _, p := range pairList[:n] {
		words = append(words, p.Key)
	}
	return words
}

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
