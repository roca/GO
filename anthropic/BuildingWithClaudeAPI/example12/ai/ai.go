package ai

import (
	"context"
	"fmt"
	"time"

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

// Add_user_message function    adds a user message to the conversation
func Add_user_message(messages []anthropic.MessageParam, message string) []anthropic.MessageParam {
	conversation := messages
	conversation = append(conversation, anthropic.NewUserMessage(anthropic.NewTextBlock(message)))
	return conversation
}

// add_assistant_message function    adds an assistant message to the conversation
func Add_assistant_message(messages []anthropic.MessageParam, message string) []anthropic.MessageParam {
	conversation := messages
	conversation = append(conversation, anthropic.NewAssistantMessage(anthropic.NewTextBlock(message)))
	return conversation
}

// chat function    sends the conversation to the API and returns the assistant's response
func Chat(conversation []anthropic.MessageParam, temperature float64, stop_sequences []string, system_options ...anthropic.TextBlockParam) (string, error) {

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
