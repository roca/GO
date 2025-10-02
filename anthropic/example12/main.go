package main

import (
	"encoding/json"
	"example12/prompt_evaluator"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"
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

	var scores []float64
	var scoreSum float64
	var evaluations []prompt_evaluator.Evaluation

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

		output := evaluator.RunTestCase(test_case)
		fmt.Printf("Test Result:\n%s\n", output)

		evaluation, err := evaluator.GradeOutput(task, output, "")
		if err != nil {
			log.Fatalf("Could not MarshalIndent evaluation output: %v", err)
		}

		evaluation.Idea = idea
		evaluation.Output = output
		evaluation.TestCase = test_case

		evaluations = append(evaluations, evaluation)

		scores = append(scores, float64(evaluation.Score))
		scoreSum = scoreSum + float64(evaluation.Score)

		fmt.Println("--------------------------------------------------")
	}

	fmt.Println("Average score:", scoreSum/float64(len(scores)))
	err = generatePromptEvaluationReport(evaluations, scores)
	if err != nil {
		log.Fatalln(err)
	}

}

func generatePromptEvaluationReport(evaluations []prompt_evaluator.Evaluation, scores []float64) error {
	outputFile, err := os.Create("output.html")
	if err != nil {
		return err
	}
	defer outputFile.Close() // Ensure the file is closed

	join := func(sep string, s []string) string {
		var s2 []string
		for _, str := range s {
			s2 = append(s2, fmt.Sprintf("<li>%s</li>", str))
		}
		return strings.Join(s2, sep)
	}

	mapJoin := func(m map[string]any) string {
		var s []string
		var unit string

		keys := make([]string, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			switch k {
			case "weight":
				unit = "kg"
			case "height":
				unit = "cm"
			default:
				unit = ""
			}
			s = append(s, fmt.Sprintf("<li><strong>%s:</strong> %v %s</li>", k, m[k], unit))
		}

		return strings.Join(s, "")
	}

	var score_sum float64

	total_tests := len(evaluations)

	for _, score := range scores {
		score_sum = score_sum + float64(score)
	}

	max_possible_score := 10
	avg_score := (score_sum / float64(len(scores))) / float64(max_possible_score) * 100

	tmpl, err := template.New("template.html").Funcs(template.FuncMap{
		"join":    join,
		"mapJoin": mapJoin,
	}).ParseFiles("template.html")
	if err != nil {
		return err
	}

	pageData := struct {
		TotalTests       int
		AvgScore         float64
		MaxPossibleScore int
		Evaluations      []prompt_evaluator.Evaluation
	}{
		TotalTests:       total_tests,
		AvgScore:         avg_score,
		MaxPossibleScore: max_possible_score,
		Evaluations:      evaluations,
	}

	err = tmpl.Execute(outputFile, pageData)
	if err != nil {
		return err
	}
	return nil
}
