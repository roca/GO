package main

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
	}
}
