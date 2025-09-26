package main

import (
	"encoding/json"
	"example12/prompt_evaluator"
	"fmt"
	"log"
)

func main() {
	evaluator := prompt_evaluator.PromptEvaluator{}
	task := prompt_evaluator.Task{
		Task: "Write a compact, concise 1 day meal plan for a single athlete",
		PromptInputSpecs: map[string]string{
			"height":       "Athlete's height in cm",
			"weight":       "Athlete's weight in kg",
			"goal":         "Goal of the athlete",
			"restrictions": "Dietery restrictions",
		},
		Format: "JSON",
	}

	ideas, err := evaluator.GenerateUniqueIdeas(
		task.Task,
		task.PromptInputSpecs,
		"dataset.json",
		3,
	)
	if err != nil {
		log.Fatalf("Could not create dataset: %v", err)
	}

	for _, idea := range ideas {
		fmt.Println("--------------------------------------------------")
		fmt.Printf("Idea: %s\n", idea)
		test_case, err := evaluator.GenerateTestCase(
			task.Task,
			idea,
			task.PromptInputSpecs,
		)
		if err != nil {
			log.Fatalf("Could not create TestCase: %v", err)
		}

		bytes, err := json.MarshalIndent(test_case, "", "\t")
		if err != nil {
			log.Fatalf("Could not MarshalIndent TestCase: %v", err)
		}

		fmt.Println(string(bytes))

		evaluator.RunTestCase(test_case)

		fmt.Println("--------------------------------------------------")
	}

}
