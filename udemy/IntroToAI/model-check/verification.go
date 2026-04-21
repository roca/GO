package main

import "fmt"

func VerifyModel(model *LoanApprovalAI, property Property, applicants []Applicant) {

	propertyName := property.Name()
	fmt.Printf("Verifying %s...\n", propertyName)

	// Run the property check on the model and the applicants
	satisfied, counterExamples := property.Check(model, applicants)

	// Print out verification result
	if satisfied {
		fmt.Printf("✅ %s is satisfied\n", propertyName)
	} else {
		fmt.Printf("❌ %s is violated. Found %d problematic cases\n",
			propertyName,
			len(counterExamples))

		// Print up to three examples for clarity
		for i := range min(3, len(counterExamples)) {

			a := counterExamples[i]
			fmt.Printf("   Example: %d, Income: $%.1fk, Credit Score: %.2f, Dept Ratio: %.2f, "+
				"Protected: %v, Decision: %v\n", i+1, a.income, a.creditScore, a.debtToIncome,
				a.protectedClass, model.ApproveLoan(a))
		}

		// If theere are more than three counterexamples, indicate that there are more
		if len(counterExamples) > 3 {
			fmt.Printf("   ... and %d more cases\n", len(counterExamples)-3)
		}
	}

}
