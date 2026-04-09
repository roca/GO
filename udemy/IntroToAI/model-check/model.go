package main

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

// Some means of determining fairness

// Check verifies if the AI model satisfies the fairness property

// Evaluate risk

// Loading the CSV file

// VerifyModel checks if the model satisfies some property.
