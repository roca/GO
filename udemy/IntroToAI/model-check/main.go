package main

import "fmt"

func main() {
	// Set up CSV file path
	csvFilePath := "loan_applicants.csv"

	// Load applicant data from CSBV

	applicants, err := LoadApplicantsFromCSV(csvFilePath)
	if err != nil {
		fmt.Printf("Error loading applicants from CSV: %v\n", err)
		return
	}

	// Define some properties we want to check (fairness and risk)
	fairnessProperty := &FairnessProperty{maxDisparity: 0.05}   // Allow up to 5% disparity in approval rates
	riskProperty := &RiskProperty{axHighRiskApprovalRate: 0.10} // Allow up to 10% approval rate for high-risk applicants

	// Create some test models
	models := []*LoanApprovalAI{
		&LoanApprovalAI{incomeWeight: 0.3, creditScoreWeight: 0.4, loanAmountWeight: 1.0, debtRatioWeight: 2.0, approvalThreshold: 5.0, employmentWeight: 0.1},
		&LoanApprovalAI{incomeWeight: 0.25, creditScoreWeight: 0.45, loanAmountWeight: 1.2, debtRatioWeight: 2.5, approvalThreshold: 4.5, employmentWeight: 0.15},
		&LoanApprovalAI{incomeWeight: 0.2, creditScoreWeight: 0.5, loanAmountWeight: 1.5, debtRatioWeight: 3.0, approvalThreshold: 4.0, employmentWeight: 0.2},
	}

	// Test each model configuration against both properties
	descriptions := []string{
		"Loan Approval AI model with Initial Parameters",
		"Loan Approval AI model with Adjusted Parameters",
		"Loan Approval AI model with Final Parameters",
	}

	for i, model := range models {
		// Print the current model's parameters
		PrintModelParams(model, descriptions[i])

		// Verify the model against both properties
		VerifyModel(model, fairnessProperty, applicants)
		VerifyModel(model, riskProperty, applicants)

	}

}
