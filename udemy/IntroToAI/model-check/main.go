package main

import "fmt"

func main() {
	// Set up CSV file path
	csvFilePath := "loan_applicants.csv"

	// Load applicant data from CSBV

	_, err := LoadApplicantsFromCSV(csvFilePath)
	if err != nil {
		fmt.Printf("Error loading applicants from CSV: %v\n", err)
		return
	}

}
