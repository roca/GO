package main

import "fmt"

type RiskProperty struct {
	axHighRiskApprovalRate float6
}

func (p *RiskProperty) Name() string {
	return "Risk Property"
}

func (p *RiskProperty) Check(model *LoanApprovalAI, applicants []Applicant) (bool, []Applicant) {
	var highRiskApproved, highRiskTotal int
	var riskyApprovals []Applicant

	// Loop through all applicants and count approvals for high-risk applicants
	for _, applicant := range applicants {
		// Make a decision
		isHoighRisk := applicant.creditScore < 0.5 && applicant.debtToIncome > 0.5

		if isHoighRisk {
			highRiskTotal++
			if model.ApproveLoan(applicant) {
				highRiskApproved++
				riskyApprovals = append(riskyApprovals, applicant)
			}
		}
	}

	if highRiskTotal == 0 {
		return true, nil // No high-risk applicants, so property is trivially satisfied
	}

	// Calculate approval rate for high-risk applicants
	highRiskApprovalRate := float64(highRiskApproved) / float64(highRiskTotal)

	// Print approval rate for debugging
	fmt.Printf("High-Risk Approval Rate: %.2f%% (Maximum allowed: %.2f%%)\n", highRiskApprovalRate*100, p.axHighRiskApprovalRate*100)

	// Check if the approval rate for high-risk applicants exceeds the maximum allowed threshold
	return highRiskApprovalRate <= p.axHighRiskApprovalRate, riskyApprovals
}
