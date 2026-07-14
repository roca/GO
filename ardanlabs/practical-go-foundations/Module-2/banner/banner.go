package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// banner("Go", 6)
// Go
// ------

func main() {
	banner("Go", 6)
	banner("G❤️", 6)

	s := "G❤️"
	fmt.Println("Len:", len(s))
	fmt.Println("s[1]:", s[1])
	fmt.Printf("s[1]: %c\n", s[1])

	for i, c := range s {
		fmt.Printf("%c at %d\n", c, i)
	}
}

func banner(text string, width int) {
	padding := (width - utf8.RuneCountInString(text)) / 2 // BUG: len is in bytes
	//padding := (width - len(text)) / 2
	fmt.Print(strings.Repeat(" ", padding))
	fmt.Println(text)
	fmt.Println(strings.Repeat("_", width))
}
