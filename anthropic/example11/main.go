package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
)

var client anthropic.Client

type Task struct {
	Task string `json:"task"`
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

	runEval(tasks)
}

func runEval(tasks []Task) {
	// Placeholder for evaluation logic
	for i, task := range tasks {
		fmt.Printf("%d: %s\n", i+1, task.Task)
	}

}

func generate_dataset() string {

	prompt := fmt.Sprintf(`
Generate a evaluation dataset for a prompt evaluation. The dataset will be used to evaluate prompts
that generate Python, JSON, or Regex specifically for AWS-related tasks. Generate an array of JSON objects,
each representing task that requires Python, JSON, or a Regex to complete.

Example output:
%sjson
[
    {
        "task": "Description of task",
    },
    ...additional
]
%s

* Focus on tasks that can be solved by writing a single Python function, a single JSON object, or a regular expression.
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
		response_stream := client.Messages.NewStreaming(context.TODO(), message_params)

		for response_stream.Next() {
			current := response_stream.Current()
			switch current.Type {
			case "content_block_delta":
				ch <- fmt.Sprintf("%s", current.Delta.Text)
			case "content_block_stop":
				close(ch)
			}
		}
	}()

	var text string

	for t := range ch {
		text = text + t
	}

	return text, nil
}
