package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/anthropics/anthropic-sdk-go"
)

var client anthropic.Client

type Task struct {
	Task   string `json:"task"`
	Format string `json:"format"`
}

func init() {
	// Install dependencies
	// Load env varaibles
	// SDK looks for 'ANTHROPIC_API_KEY' env by default

	// Create an API Client
	client = anthropic.NewClient()
}

func main() {
	prompt := generate_dataset()

	var conversation []anthropic.MessageParam

	conversation = add_user_message(conversation, prompt)
	conversation = add_assistant_message(conversation, "```json")

	var stop_sequences []string

	stop_sequences = append(stop_sequences, "```")

	text, err := chat(conversation, 0.0, stop_sequences)
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

	results, err := runEval(tasks)
	if err != nil {
		log.Fatal(err)
	}

	for i, result := range results {
		fmt.Printf("(%s) Task %d. %s\n", result.TestCase.Format, i+1, result.TestCase.Task)
		fmt.Printf("Score: %.2f\n", result.Score)
		fmt.Printf("Reasoning: %s\n", result.Reasoning)
		fmt.Printf("%s\n-----------------------------------------------------------\n\n", result.Output)
	}

	//Average score
	var total_score float64
	for _, result := range results {
		total_score += result.Score
	}
	average_score := total_score / float64(len(results))
	fmt.Printf("Average Score: %.2f\n", average_score)

}

// validateJSON function    Validates if the input is a valid JSON
func validateJSON(text string) (int, error) {
	// Unmarshal the text into a map
	var js map[string]interface{}
	err := json.Unmarshal([]byte(text), &js)
	if err != nil {
		return 0, err
	}
	return 10, nil
}

// validateGo function    Validates if the input is a valid Go code
func validateGo(text string) (int, error) {
	// Create a new FileSet. This is required by the parser to manage source file positions.
	fset := token.NewFileSet()

	// Parse the text and ceate an AST
	_, err := parser.ParseFile(fset, "", text, parser.AllErrors)
	if err != nil {
		return 0, err
	}
	return 10, nil
}

// validateRegex function    Validates if the input is a valid Regex
func validateRegex(text string) (int, error) {
	// Compile the regex
	_, err := regexp.Compile(text)
	if err != nil {
		return 0, err
	}
	return 10, nil
}

// runTestCase function    Merges the prompt and the test case input and then returns the result
func runPrompt(test_case Task) (string, error) {

	prompt := fmt.Sprintf(`
		Please solve the following task:

		%s

		* Respond only with Go, JSON, or a plain regex, depending on what the task requires.
		* Do not add any comments or commentary or explanation.
		* Do not use any pyhton in coding examples.
`, test_case.Task)

	var conversation []anthropic.MessageParam
	var stop_sequences []string

	conversation = add_user_message(conversation, prompt)
	conversation = add_assistant_message(conversation, fmt.Sprintf("```%s", test_case.Format))
	text, err := chat(conversation, 0.0, stop_sequences)
	if err != nil {
		return "", err
	}

	// fmt.Println(text)

	return text, nil
}

type Result struct {
	Output    string
	TestCase  Task
	Score     float64
	Reasoning string
}

type Evaluation struct {
	Strengths  []string `json:"strengths"`
	Weaknesses []string `json:"weaknesses"`
	Reasoning  string   `json:"reasoning"`
	Score      float64  `json:"score"`
}

func gradeByModel(test_case Task, output string) (Evaluation, error) {

	eval_prompt := fmt.Sprintf(`
You are an expert code reviewer. Evaluate this AI-generated solution.
    
Task: %s
Solution: %s
    
Provide your evaluation as a structured JSON object with:
- "strengths": An array of 1-3 key strengths
- "weaknesses": An array of 1-3 key areas for improvement  
- "reasoning": A concise explanation of your assessment
- "score": A number between 1-10
`, test_case.Task, output)

	var conversation []anthropic.MessageParam
	var stop_sequences []string

	stop_sequences = append(stop_sequences, "```")

	conversation = add_user_message(conversation, eval_prompt)
	conversation = add_assistant_message(conversation, "```json")
	text, err := chat(conversation, 0.0, stop_sequences)
	if err != nil {
		return Evaluation{}, err
	}

	var evaluation Evaluation

	err = json.Unmarshal([]byte(text), &evaluation)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return Evaluation{}, err
	}

	// fmt.Println(text)

	return evaluation, nil
}

// runTestCase function    Calls runPrompt and then grades the result
func runTestCase(test_case Task) (*Result, error) {
	output, err := runPrompt(test_case)
	if err != nil {
		return nil, err
	}

	evaluation, err := gradeByModel(test_case, output)
	if err != nil {
		return nil, err
	}

	result := Result{
		Output:    output,
		TestCase:  test_case,
		Score:     evaluation.Score,
		Reasoning: evaluation.Reasoning,
	}

	return &result, nil
}

// runEval function    Loads the dataset and calls runTestCase with each case
func runEval(tasks []Task) ([]Result, error) {

	var results []Result

	// Placeholder for evaluation logic
	for _, task := range tasks {
		// fmt.Printf("Task(%d): %s\n", i+1, task.Task)
		result, err := runTestCase(task)
		if err != nil {
			return nil, err
		}

		results = append(results, *result)
	}

	return results, nil
}

func generate_dataset() string {

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

// add_user_message function    adds a user message to the conversation
func add_user_message(messages []anthropic.MessageParam, message string) []anthropic.MessageParam {
	conversation := messages
	conversation = append(conversation, anthropic.NewUserMessage(anthropic.NewTextBlock(message)))
	return conversation
}

// add_assistant_message function    adds an assistant message to the conversation
func add_assistant_message(messages []anthropic.MessageParam, message string) []anthropic.MessageParam {
	conversation := messages
	conversation = append(conversation, anthropic.NewAssistantMessage(anthropic.NewTextBlock(message)))
	return conversation
}

// chat function    sends the conversation to the API and returns the assistant's response
func chat(conversation []anthropic.MessageParam, temperature float64, stop_sequences []string, system_options ...anthropic.TextBlockParam) (string, error) {

	//	incomming_message := conversation[len(conversation)-1]
	//
	//	fmt.Printf("%s\n", *incomming_message.Content[0].GetText())
	//	fmt.Println("----------------------------")

	message_params := anthropic.MessageNewParams{
		MaxTokens:     1024,
		Model:         anthropic.ModelClaude3_5HaikuLatest,
		Messages:      conversation,
		Temperature:   anthropic.Float(temperature),
		StopSequences: stop_sequences,
	}

	if len(system_options) > 0 {
		message_params.System = system_options // Optional system-level instructions to the model. (e.g. 'You are a helpful assistant.')
	}

	ch := make(chan string, 1)

	go func() {
		// Make a new Request

		timeout_ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Minute))
		defer func() {
			cancel()
			// fmt.Println("context cancelled")
			close(ch)
		}()

		response_stream := client.Messages.NewStreaming(timeout_ctx, message_params)

		for response_stream.Next() {
			current := response_stream.Current()
			switch current.Type {
			case "content_block_delta":
				ch <- fmt.Sprintf("%s", current.Delta.Text)
			case "content_block_stop":
				return
			}
		}
	}()

	var text string

	for t := range ch {
		text = text + t
	}

	return text, nil
}
