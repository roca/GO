package files

import (
	"bufio"
	_ "embed"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"github.com/rsc/pdf"
)

// var (
// 	//go:embed iris.csv
// 	irisCsv string
// )

func OpenTextFile(filePath string) {
	// Method: 1
	fmt.Println("Method1: Using 'ioutil' package")
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	// Method: 2
	fmt.Println("\nMethod2: Using 'os' package")
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
func OpenCSVFile(filePath string) {
	// Open the File
	csvfile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer csvfile.Close()

	// // Method: 1
	// fmt.Println("Method1: Using 'csv' package")
	// r := csv.NewReader(csvfile)
	// for {
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(record)
	// }

	// // Method: 2
	// fmt.Println("Method1: Using Open CSV with non Standard lib package")

	// // Method: 3
	// fmt.Println("\nMethods3: Using 'embed' package")
	// fmt.Println(irisCsv)

	// Method 4
	df := dataframe.ReadCSV(csvfile)
	fmt.Println(df)
}

func OpenPDFFile(filePath string) {
	// Open the File
	pdfFile, err := pdf.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	//pdfFile1, err := must(pdf.Open,filePath)
	fmt.Println(pdfFile.Page(1).Content())
}
func must(f func(inputs ...interface{}) (interface{}, error), inputs ...interface{}) (interface{}, error) {
	result, err := f(inputs)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}
