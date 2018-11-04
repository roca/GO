// File name: ...\s15\03_io1\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type team []string

var fileName = "names.txt"
var newFileName = "new.txt"

func main() {

	// uncomment one at a time.
	writeToFile()
	// readFromFile()
	// renameFile()
	// removeFile()
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
func writeToFile() {
	players := team{"Messi", "Pele", "Maradona", "Zidane", "Platini"}

	err := ioutil.WriteFile(fileName, []byte(players.toString()), 0666)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Wrote the following to file: %v \n", players.toString())
	}
}

func (t team) toString() string {
	return strings.Join([]string(t), ",")
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
func readFromFile() {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
	}

	s := strings.Split(string(bs), ",")
	fmt.Println(s)
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
func renameFile() {
	err := os.Rename(fileName, newFileName)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File renamed.")
	}
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
func removeFile() {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File removed.")
	}
}
