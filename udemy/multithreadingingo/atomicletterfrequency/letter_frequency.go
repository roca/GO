package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency *[26]int32) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		index := strings.Index(allLetters, c)
		if index >= 0 {
			frequency[index] += 1
		}
	}
}

func main() {
	var frequency [26]int32
	start := time.Now()
	for i := 1000; i <= 1200; i++ {
		go countLetters(fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequency)
	}
	elapsed := time.Since(start)
	fmt.Printf("Processing took %s\n", elapsed)
	fmt.Println("Done")
	for i, f := range frequency {
		fmt.Printf("%s: %d\n", string(allLetters[i]), f)
	}
}
