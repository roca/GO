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

var system_prompt_test_cases = anthropic.TextBlockParam{
	Text: `
You are a test case creator specializing in designing evaluation scenarios.
	`,
}
var system_prompt_ideas = anthropic.TextBlockParam{
	Text: `
You are a test scenario designer specialized in creating diverse, unique testing scenarios.
	`,
}

type Idea string

type PromptEvaluator struct {
}

func (pe *PromptEvaluator) GenerateTestCase(
	task_description string,
	idea Idea,
	prompt_inputs_spec map[string]string,
) {

	prompt := generateTestCasePrompt(task_description, idea, prompt_inputs_spec)
	// fmt.Println(prompt)

	var conversation []anthropic.MessageParam

	conversation = ai.Add_user_message(conversation, prompt)
	conversation = ai.Add_assistant_message(conversation, "```json")

	var stop_sequences []string

	stop_sequences = append(stop_sequences, "```")
	text, err := ai.Chat(conversation, 0.7, stop_sequences, system_prompt_test_cases)
	if err != nil {
		log.Fatalf("Chat error: %v", err)
	}

	fmt.Println(text)

}

func (pe *PromptEvaluator) GenerateUniqueIdeas(
	task_description string,
	prompt_inputs_spec map[string]string,
	output_file string,
	num_cases int,
) ([]Idea, error) {

	prompt := generateIdeasPrompt(task_description, prompt_inputs_spec, num_cases)
	// fmt.Println(prompt)

	var conversation []anthropic.MessageParam

	conversation = ai.Add_user_message(conversation, prompt)
	conversation = ai.Add_assistant_message(conversation, "```json")

	var stop_sequences []string

	stop_sequences = append(stop_sequences, "```")

	text, err := ai.Chat(conversation, 1.0, stop_sequences, system_prompt_ideas)
	if err != nil {
		return []Idea{}, err
	}

	err = os.WriteFile(output_file, []byte(text), 0644) // 0644 sets file permissions
	if err != nil {
		return []Idea{}, err
	}
	log.Println("File written successfully.")

	// read from dataset.json

	bytes, err := os.ReadFile(output_file)
	if err != nil {
		return []Idea{}, err
	}

	var ideas []Idea

	err = json.Unmarshal(bytes, &ideas)
	if err != nil {
		return []Idea{}, err
	}

	return ideas, nil
}

func generateIdeasPrompt(
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
func generateTestCasePrompt(
	task_description string,
	idea Idea,
	prompt_inputs_spec map[string]string,
) string {

	allowed_keys := []string{}
	example_prompt_inputs := []string{}

	for k, v := range prompt_inputs_spec {
		allowed_keys = append(allowed_keys, fmt.Sprintf(`"%s"`, k))
		example_prompt_inputs = append(example_prompt_inputs, fmt.Sprintf(`"%s": "EXAMPLE_VALUE", // %s`, k, v))
	}

	prompt := fmt.Sprintf(` 
        Generate a single detailed test case for a prompt evaluation based on:
        
        <task_description>
        %s
        </task_description>
        
        <specific_idea>
        %s
        </specific_idea>
        
        <allowed_input_keys>
        %s
        </allowed_input_keys>
        
        Output Format:
        %s
        %sjson
        {{
            "prompt_inputs": {{
            		%s
            }},
            "solution_criteria": ["criterion 1", "criterion 2", ...] // Concise list of criteria for evaluating the solution, 1 to 4 items
        }}
        %s
        
        IMPORTANT REQUIREMENTS:
        - You MUST ONLY use these exact input keys in your prompt_inputs: {allowed_keys}        
        - Do NOT add any additional keys to prompt_inputs
        - All keys listed in allowed_input_keys must be included in your response
        - Make the test case realistic and practically useful
        - Include measurable, concise solution criteria
        - The solution criteria should ONLY address the direct requirements of the task description and the generated prompt_inputs
        - Avoid over-specifying criteria with requirements that go beyond the core task
        - Keep solution criteria simple, focused, and directly tied to the fundamental task
        - The test case should be tailored to the specific idea provided
        - Quick to solve without requiring extensive computation or multi-step processing
        - Solvable with no more than 400 tokens of output
        - DO NOT include any fields beyond those specified in the output format

        Here's an example of a sample input with an ideal output:
        <sample_input>
        <sample_task_description>
        Extract topics out of a passage of text
        </sample_task_description>
        <sample_specific_idea>
        Testing with a text that contains multiple nested topics and subtopics (e.g., a passage about renewable energy that covers solar power economics, wind turbine technology, and policy implications simultaneously)
        </sample_specific_idea>

        <sample_allowed_input_keys>
        "content"
        </sample_allowed_input_keys>
        </sample_input>
        <ideal_output>
        %sjson
        {
            "prompt_inputs": {
                "content": "The transition to renewable energy encompasses numerous interdependent dimensions. Solar photovoltaic technology has seen dramatic cost reductions, with panel efficiency improving 24 percent since 2010 while manufacturing costs declined by 89 percent, making it economically competitive with fossil fuels in many markets. Concurrently, wind energy has evolved through innovative turbine designs featuring carbon-fiber composite blades and advanced control systems that increase energy capture by 35 percent in low-wind conditions."
            },
            "solution_criteria": [
                "Includes all topics mentioned"   
            ]
        }
        %s
        </ideal_output>
        This is ideal output because the solution criteria is concise and doesn't ask for anything outside of the scope of the task description.
        `, task_description, idea, strings.Join(allowed_keys, ", "), strings.Join(example_prompt_inputs, "\n\t"), "```", "ppp", "```", "```", "```")

	return prompt
}
