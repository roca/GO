package main

import (
	"example12/prompt_evaluator"
	"log"
)

func main() {
	evaluator := prompt_evaluator.PromptEvaluator{}

	_, err := evaluator.GenerateDataset(
		"Write a compact, concise 1 day meal plan for a single athlete",
		map[string]string{
			"height":       "Athlete's height in cm",
			"weight":       "Athlete's weight in kg",
			"goal":         "Goal of the athlete",
			"restrictions": "Dietery restrictions",
		},
		"dataset.json",
		3,
	)
	if err != nil {
		log.Fatalf("Could not create dataset: %v", err)
	}

}
