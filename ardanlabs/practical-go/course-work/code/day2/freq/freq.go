package main

// Q: What is the most common word (ignoring case) in sherlock.txt?

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
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

	wordCount, err := wordFrequency(file)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
	// count the frequency of each word
	// print the most common word
}

func wordFrequency(r io.Reader) (map[string]int, error) {
	// read from r one line at a time
	s := bufio.NewScanner(r)
	wordCount := make(map[string]int)
	lnum := 0
	for s.Scan() {
		words := wordRe.FindAllString(s.Text(), -1) // current line
		if len(words) != 0 {
			lnum++
			fmt.Println(words)
			for _, word := range words {
				wordCount[word]++
			}
			break
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	fmt.Println("num lines:", lnum)
	// split the text into words
	// count the frequency of each word
	// return the frequency map

	return wordCount, nil
}
