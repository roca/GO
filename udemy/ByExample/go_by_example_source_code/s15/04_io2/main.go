// File name: ...\s15\04_io2\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// How to run? go run main.go inp.txt arg2 arg3

func main() {

	// uncomment one at a time.
	readFile()
	// readWrods()
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
func readFile() {
	f, err := os.Open(os.Args[1])

	for i, v := range os.Args {
		fmt.Println(i, v)
	}

	fmt.Println()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	} else {
		io.Copy(os.Stdout, f)
	}

	currentDir()
}

func currentDir() {
	var err error
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	fmt.Println("\nCurrent directory=", cwd)
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
func readWrods() {
	input := bufio.NewScanner(os.Stdin)

	var all string
	fmt.Println("Enter some words. To exist, type 'end'.")

	for input.Scan() {
		s := input.Text()

		if s == "end" {
			break
		}

		all += (s + " ")
	}

	fmt.Println(all)
}
