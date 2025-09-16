package main

import (
	"context"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
)

var client anthropic.Client

func init() {
	// Install dependencies
	// Load env varaibles
	// SDK looks for 'ANTHROPIC_API_KEY' env by default

	// Create an API Client
	client = anthropic.NewClient()
}

func main() {
	var conversation []anthropic.MessageParam

	userInput := "Generate  a very short event bridge rule as json"
	conversation = add_user_message(conversation, userInput)

	assistant_prefilled_message := "```json"
	conversation = add_assistant_message(conversation, assistant_prefilled_message)

	err := chat(conversation, 1.0, []string{"```"})
	if err != nil {
		panic(err.Error())
	}

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
func chat(conversation []anthropic.MessageParam, temperature float64, stop_sequence []string, system_options ...anthropic.TextBlockParam) error {

	incomming_message := conversation[len(conversation)-1]

	fmt.Printf("%s\n", *incomming_message.Content[0].GetText())
	fmt.Println("----------------------------")

	message_params := anthropic.MessageNewParams{
		MaxTokens:     1024,
		Model:         anthropic.ModelClaude3_5HaikuLatest,
		Messages:      conversation,
		Temperature:   anthropic.Float(temperature),
		StopSequences: stop_sequence,
	}

	if len(system_options) > 0 {
		message_params.System = system_options // Optional system-level instructions to the model. (e.g. 'You are a helpful assistant.')
	}

	// Make a new Request
	response_stream := client.Messages.NewStreaming(context.TODO(), message_params)

	ch := make(chan string, 1)

	go func() {
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

	for text := range ch {
		fmt.Printf("%s", text)
	}

	fmt.Println()

	return nil
}
