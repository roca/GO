package prompt_evaluator

import (
	"encoding/json"
	"example12/ai"
	"fmt"
	"log"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
)

type Task struct{}

type PromptEvaluator struct {
}

func (pe *PromptEvaluator) GenerateDataset(
	task_description string,
	prompt_inputs_spec map[string]string,
	output_file string,
	num_cases int,
) ([]Task, error) {

	prompt := generatePrompt(task_description, prompt_inputs_spec, output_file, num_cases)

	var conversation []anthropic.MessageParam

	conversation = ai.Add_user_message(conversation, prompt)
	conversation = ai.Add_assistant_message(conversation, "```json")

	var stop_sequences []string

	stop_sequences = append(stop_sequences, "```")

	text, err := ai.Chat(conversation, 0.0, stop_sequences)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("dataset.json", []byte(text), 0644) // 0644 sets file permissions
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File written successfully.")

	// read from dataset.json

	bytes, err := os.ReadFile("dataset.json")
	if err != nil {
		log.Fatal(err)
	}

	var tasks []Task

	err = json.Unmarshal(bytes, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	return []Task{}, nil

}

func generatePrompt(
	task_description string,
	prompt_inputs_spec map[string]string,
	output_file string,
	num_cases int,
) string {

	prompt := fmt.Sprintf(`
Generate a evaluation dataset for a prompt evaluation. The dataset will be used to evaluate prompts
that generate Go, JSON, or Regex specifically for AWS-related tasks. Generate an array of JSON objects,
each representing task and format that requires Go, JSON, or a Regex to complete.

Example output:
%sjson
[
    {
        "task": "Description of task",
				"format": "go" or "json" or "regex"
    },
    ...additional
]
%s

* Focus on tasks that can be solved by writing a single Go function, a single JSON object, or a regular expression.
* Focus on tasks that do not require writing much code

Please generate 3 objects.
`, "```", "```")

	return prompt
}
