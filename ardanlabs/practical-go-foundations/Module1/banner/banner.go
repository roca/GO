package main

import (
	"fmt"
	"strings"
)

// banner("Go", 6)
// Go
// ------

func main() {
	banner("Go", 6)
	banner("G❤️", 6)

	s := "G❤️"
	fmt.Println("Len:", len(s))
	fmt.Println("s[0]:", s[0])
	fmt.Println("s[1]:", s[1])
}

func banner(text string, width int) {
	padding := (width - len(text)) / 2
	fmt.Print(strings.Repeat(" ", padding))
	fmt.Println(text)
	fmt.Println(strings.Repeat("_", width))
}
