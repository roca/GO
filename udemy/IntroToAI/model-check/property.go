package main

import "fmt"

type Property interface {
	Check(model *LoanApprovalAI, applicants []Applicant) (bool, []Applicant)
	Name() string
}

type FairnessProperty struct {
	maxDisparity float64
}

func (p *FairnessProperty) Name() string {
	return "Fairness Property"
}

func (p *FairnessProperty) Check(model *LoanApprovalAI, applicants []Applicant) (bool, []Applicant) {
	var protectedApproved, protectedTotal, nonProtectedApproved, nonProtectedTotal int
	var unfairDecisions []Applicant

	// Loop through all applicants and count approvals for each group
	for _, applicant := range applicants {
		// Make a decision
		decision := model.ApproveLoan(applicant)

		// Update counters based on protected status class
		if applicant.protectedClass {
			protectedTotal++
			if decision {
				protectedApproved++
			}
		} else {
			nonProtectedTotal++
			if decision {
				nonProtectedApproved++
			}
		}

		// Check for potentially problenatic individual decisions.
		if !decision &&
			applicant.creditScore > 0.7 &&
			applicant.debtToIncome < 0.3 &&
			applicant.income > 60 {
			unfairDecisions = append(unfairDecisions, applicant)
		}
	}

	// Calculate approval rates for each group.
	protectedRate := float64(protectedApproved) / float64(protectedTotal)
	nonProtectedRate := float64(nonProtectedApproved) / float64(nonProtectedTotal)

	// Calculate maxDisparity
	disparity := nonProtectedRate - protectedRate // How much more likely non-protected applicants are to be approved compared to protected applicants

	// Print approval rates and disparity for debugging
	fmt.Printf("Approval Rate - Protected: %.2f%%, Non-Protected: %2f%%\n", protectedRate*100, nonProtectedRate*100)
	fmt.Printf("Disparity: %2f%% (Maximum allowed: %2f%%)\n", disparity*100, p.maxDisparity*100)

	// Check if the disparity exceeds the maximum allowed threshold
	isFair := disparity <= p.maxDisparity && len(unfairDecisions) == 0

	return isFair, unfairDecisions
}
