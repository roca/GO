package main

import (
	"fmt"
	"os"

	"github.com/roca/GO/udemy/DataScienceNLP/files"
)

func main() {
	stringExamples()
}

func stringExamples() {
	fmt.Println("Method1: Create Characters with type ")
	var char byte = 'A'
	var char2 rune = 'A'
	fmt.Printf("Char as Byte %d %T\n", char, char)
	fmt.Printf("Char as Rune %d %T\n", char2,char2)

	fmt.Println("Method2: Create Characters with Method")
	charA := byte('A')
	charB := rune('A')
	fmt.Printf("Char as Byte:Fxn %d %T\n", charA, charA)
	fmt.Printf("Char as Rune:Fxn %d %T\n", charB,charB)

	fmt.Println("Actual Representation")
	str1 := string(char)
	str2 := string(char2)
	fmt.Printf("Char as String %s %T\n", str1, str1)
	fmt.Printf("Char as String %s %T\n", str2,str2)

	fmt.Println("Representation method 2 using Printf")
	fmt.Printf("Char as Byte %c %T\n", char, char)
	fmt.Printf("Char as Rune %c %T\n", char2,char2)
}
func filesExmaples() {
	args := os.Args[1:]
	if len(args) == 2 {
		switch args[1] {
		case "csv":
			files.OpenCSVFile(args[0])
		case "pdf":
			files.OpenPDFFile(args[0])
		case "txt":
			files.OpenTextFile(args[0])
		default:
			fmt.Println("Usage: file type[csv|txt|pdf]")
		}
	} else {
		fmt.Print("Usage: file type[csv|txt|pdf]")
	}
}
