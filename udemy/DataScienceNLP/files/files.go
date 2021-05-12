package files

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

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

	// Method: 1
	fmt.Println("Method1: Using 'csv' package")
	r := csv.NewReader(csvfile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}

	// Method: 2
	fmt.Println("Method1: Using Open CSV with non Standard lib package")

}
