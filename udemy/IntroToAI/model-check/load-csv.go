package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func LoadApplicantsFromCSV(filePath string) ([]Applicant, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read the header row
	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("could not read header: %v", err)
	}

	// Map header columns to struct fields
	columnIndices := map[string]int{
		"income":         -1,
		"creditScore":    -1,
		"loanAmount":     -1,
		"debtToIncome":   -1,
		"yearsEmployed":  -1,
		"protectedClass": -1,
	}

	// Find the column indices
	for i, column := range header {
		col := strings.ToLower(strings.TrimSpace(column))
		switch {
		case strings.Contains(col, "income") && !strings.Contains(col, "debt"):
			columnIndices["income"] = i
		case strings.Contains(col, "credit"):
			columnIndices["creditScore"] = i
		case strings.Contains(col, "loan"):
			columnIndices["loanAmount"] = i
		case strings.Contains(col, "debt"):
			columnIndices["debtToIncome"] = i
		case strings.Contains(col, "employ"):
			columnIndices["yearsEmployed"] = i
		case strings.Contains(col, "protect"):
			columnIndices["protectedClass"] = i
		}
	}

	// Verify that all required columns were found
	for field, idx := range columnIndices {
		if idx == -1 {
			return nil, fmt.Errorf("missing required column for field: %s in CSV", field)
		}
	}

	// Read applicant data
	var applicants []Applicant
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read records: %v", err)
	}

	for i, record := range records {
		// Parse values
		income, err := parseFloat(record[columnIndices["income"]])
		if err != nil {
			return nil, fmt.Errorf("error parsing income on row %d: %v", i+2, err)
		}
		// convert income to thousands
		income /= 1000

		creditScore, err := parseFloat(record[columnIndices["creditScore"]])
		if err != nil {
			return nil, fmt.Errorf("error parsing credit score on row %d: %v", i+2, err)
		}
		// Normalize credit score if in 300 - 850 range
		if creditScore > 1 {
			creditScore = (creditScore - 300) / (850 - 300)
		}

		loanAmount, err := parseFloat(record[columnIndices["loanAmount"]])
		if err != nil {
			return nil, fmt.Errorf("error parsing loan amount on row %d: %v", i+2, err)
		}
		// convert loan amount to thousands
		loanAmount /= 1000

		debtToIncome, err := parseFloat(record[columnIndices["debtToIncome"]])
		if err != nil {
			return nil, fmt.Errorf("error parsing debt-to-income ratio on row %d: %v", i+2, err)
		}
		// Normalize debt-to-income ratio if in percentage format
		if debtToIncome > 1 {
			debtToIncome /= 100
		}

		yearsEmployed, err := parseFloat(record[columnIndices["yearsEmployed"]])
		if err != nil {
			return nil, fmt.Errorf("error parsing years employed on row %d: %v", i+2, err)
		}

		protectedClass, err := parseBool(record[columnIndices["protectedClass"]])
		if err != nil {
			return nil, fmt.Errorf("error parsing protected class on row %d: %v", i+2, err)
		}

		applicant := Applicant{
			income:         income,
			creditScore:    creditScore,
			loanAmount:     loanAmount,
			debtToIncome:   debtToIncome,
			yearsEmployed:  yearsEmployed,
			protectedClass: protectedClass,
		}

		applicants = append(applicants, applicant)
	}

	fmt.Printf("Successfully loaded %d applicants from CSV.\n", len(applicants))

	return applicants, nil

}
