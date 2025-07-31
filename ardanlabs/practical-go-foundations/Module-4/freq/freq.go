package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
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

	for scanner.Scan() {
		words := wordRe.FindAllString(scanner.Text(),-1)
	}

	scanner.Split(bufio.ScanWords)

	// Set the scanner to split by words
	scanner.Split(bufio.ScanWords)

	wordFrequencies := make(map[string]int)

	for scanner.Scan() {
		word := scanner.Text()
		wordFrequencies[word]++
		//fmt.Println(word)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %v", err)
	}

	pairList := rankByWordCount(wordFrequencies)

	fmt.Println(pairList[:5])
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
