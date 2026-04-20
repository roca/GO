package main

import "fmt"

type LoanApprovalAI struct {
	// Weights for different factors used in the decision making process
	incomeWeight      float64
	creditScoreWeight float64
	loanAmountWeight  float64
	debtRatioWeight   float64
	employmentWeight  float64
	approvalThreshold float64
}

type Applicant struct {
	income         float64 // Annual income in thousands, so 50 = $50,000.00
	creditScore    float64 // Credit score normalized to 0-1, from a typical 300-850 range.
	loanAmount     float64 // Requested loan amount in thousands, so 20 = $20,000.00
	debtToIncome   float64 // debt to income ration (0-1), already normalized
	yearsEmployed  float64
	protectedClass bool // Whether or not the applicant belongs to some protected class
}

// ApproveLoan determines if the applicant shouild be approved for a loan.
func (ai *LoanApprovalAI) ApproveLoan(applicant Applicant) bool {
	loanToIncomeRatio := 0.0
	if applicant.income > 0 {
		loanToIncomeRatio = applicant.loanAmount / applicant.income
	}

	score := applicant.income*ai.incomeWeight +
		applicant.creditScore*ai.creditScoreWeight -
		loanToIncomeRatio*ai.loanAmountWeight -
		applicant.debtToIncome*ai.debtRatioWeight +
		applicant.yearsEmployed*ai.employmentWeight

	return score > ai.approvalThreshold
}

func PrintModelParams(model *LoanApprovalAI, description string) {
	fmt.Printf("\n===== %s =====\n", description)
	fmt.Printf("- Income Weight: %.2f\n", model.incomeWeight)
	fmt.Printf("- Credit Score Weight: %.2f\n", model.creditScoreWeight)
	fmt.Printf("- Loan Amount Weight: %.2f\n", model.loanAmountWeight)
	fmt.Printf("- Debt Ratio Weight: %.2f\n", model.debtRatioWeight)
	fmt.Printf("- Employment Weight: %.2f\n", model.employmentWeight)
	fmt.Printf("- Approval Threshold: %.2f\n", model.approvalThreshold)
}

// Some means of determining fairness

// Check verifies if the AI model satisfies the fairness property

// Evaluate risk

// Loading the CSV file

// VerifyModel checks if the model satisfies some property.
