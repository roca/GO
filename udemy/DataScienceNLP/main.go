package main

import (
	"fmt"
	"os"

	"github.com/roca/GO/udemy/DataScienceNLP/files"
)

func main() {
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
