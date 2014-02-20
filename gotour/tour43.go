// tour43
package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) (word_map map[string]int) {

	words := strings.Split(s, " ")
	word_map = make(map[string]int)

	for _, value := range words {

		_, ok := word_map[value]

		if ok {
			word_map[value]++
		} else {
			word_map[value] = 1
		}
	}

	return
}

func main() {
	wc.Test(WordCount)
}
