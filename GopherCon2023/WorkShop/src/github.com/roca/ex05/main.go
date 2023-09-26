package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	bytes, err := os.ReadFile("proverbs.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")
	i := 1
	for _, l := range lines {
		fmt.Printf("%d %s (WC: %d)\n", i, l, len(strings.Fields(l)))
		for k, v := range charCount(l) {
			fmt.Printf("'%s'=%d, ", k, v)
		}
		fmt.Print("\n\n")
		i++
	}
}

func charCount(line string) map[string]int {
	m := make(map[string]int, 0)
	for _, char := range line {
		m[string(char)] = m[string(char)] + 1
	}
	return m
}
