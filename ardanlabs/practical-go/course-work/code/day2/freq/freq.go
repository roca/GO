package main

// Q: What is the most common word (ignoring case) in sherlock.txt?

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// read the file
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	wordFrequency(file)
	// count the frequency of each word
	// print the most common word
}

func wordFrequency(r io.Reader) (map[string]int, error) {
	// read from r one line at a time
	s := bufio.NewScanner(r)
	lnum := 0
	for s.Scan() {
		lnum++
		s.Text() // current line
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	fmt.Println("num lines:", lnum)
	// split the text into words
	// count the frequency of each word
	// return the frequency map

	return nil, nil
}
