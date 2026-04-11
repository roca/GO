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

	// Define some properties we want to check (fairness and risk)

	// Create some test models

	// Test each model configuration against both properties

}
