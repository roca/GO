package prompt_evaluator

import (
	"encoding/json"
	"example12/ai"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
)

var system_prompt = anthropic.TextBlockParam{
	Text: `
You are a test case creator specializing in designing evaluation scenarios.
	`,
}

type Task string

type PromptEvaluator struct {
}

func (pe *PromptEvaluator) GenerateDataset(
	task_description string,
	prompt_inputs_spec map[string]string,
	output_file string,
	num_cases int,
) ([]Task, error) {

	prompt := generatePrompt(task_description, prompt_inputs_spec, num_cases)
	fmt.Println(prompt)

	var conversation []anthropic.MessageParam

	conversation = ai.Add_user_message(conversation, prompt)
	conversation = ai.Add_assistant_message(conversation, "```json")

	var stop_sequences []string

	stop_sequences = append(stop_sequences, "```")

	text, err := ai.Chat(conversation, 0.0, stop_sequences, system_prompt)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(output_file, []byte(text), 0644) // 0644 sets file permissions
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File written successfully.")

	// read from dataset.json

	bytes, err := os.ReadFile(output_file)
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
	num_cases int,
) string {

	prompt_inputs := []string{}

	for k, v := range prompt_inputs_spec {
		prompt_inputs = append(prompt_inputs, fmt.Sprintf(`"%s": "%s"`, k, v))
	}

	prompt := fmt.Sprintf(`
				Generate %d unique, diverse ideas for testing a prompt that accomplishes this task:
        
        <task_description>
        %s
        </task_description>

        The prompt will receive the following inputs
        <prompt_inputs>
        %s
        </prompt_inputs>
        
        Each idea should represent a distinct scenario or example that tests different aspects of the task.
        
        Output Format:
        Provide your response as a structured JSON array where each item is a brief description of the idea.
        
        Example:
        %sjson
        [
            "Testing with technical computer science terminology",
            "Testing with medical research findings",
            "Testing with complex mathematical concepts",
            ...
        ]
        %s
        
        Ensure each idea is:
        - Clearly distinct from the others
        - Relevant to the task description
        - Specific enough to guide generation of a full test case
        - Quick to solve without requiring extensive computation or multi-step processing
        - Solvable with no more than 400 tokens of output

        Remember, only generate %d unique ideas
`, num_cases, task_description, strings.Join(prompt_inputs, "\n\t"), "```", "```", num_cases)

	return prompt
}
