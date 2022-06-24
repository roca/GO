package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var (
	matches []string
)

func fileSearch(root string, filename string) {
	fmt.Println("Searching in", root)
	files, _ := ioutil.ReadDir(root)
	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			matches = append(matches, filepath.Join(root, file.Name()))
		}
		if file.IsDir() {
			fileSearch(filepath.Join(root, file.Name()), filename)
		}
	}
}

func main() {
	fileSearch("/Users/romelcampbell/GOCODE/udemy", "README.md")
	for _, file := range matches {
		fmt.Println("Matched", file)
	}
}
