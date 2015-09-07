package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tealeg/xlsx"
)

func usage(arg string) string {
	return fmt.Sprintf("usage: %s <path to excel File>\n", arg)
}

func main() {

	if len(os.Args) == 1 {
		fmt.Printf(usage(filepath.Base(os.Args[0])))
		os.Exit(1)
	}
	excelFileName := os.Args[len(os.Args)-1]

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("Error")
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				fmt.Printf("%s\n", cell.String())
			}
		}
	}
}
