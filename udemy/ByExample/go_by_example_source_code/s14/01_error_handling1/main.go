// File name: ...\s14\01_error_handling1\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"log"
	"os"
)

var fileName = "non_existing.txt"

func main() {

	// Run test cases one at a time.

	testCase1()
	fmt.Println("Control has returned to main after an error in testCase1().")

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	// testCase2()
	// fmt.Println("Control does not return to main after an error in testCase2().")

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	// testCase3()
	// fmt.Println("Control does not return to main after panic in testCase3().")
}

func testCase1() {
	_, err := os.Open(fileName)
	if err != nil {
		log.Println("Error: ", err)
	}
}

func testCase2() {
	_, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
}

func testCase3() {
	_, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
}
