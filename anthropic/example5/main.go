package main

import (
	"context"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
)

var client anthropic.Client
var system_prompt = anthropic.TextBlockParam{
	Text: `
	You are a python software engineer who writes very concise code.
	Respond with answers as concisely as possible.
	`,
}

func init() {
	// Install dependencies
	// Load env varaibles
	// SDK looks for 'ANTHROPIC_API_KEY' env by default

	// Create an API Client
	client = anthropic.NewClient()
}

func main() {
	var conversation []anthropic.MessageParam

	userInput := "Write a Python function that checks a string for duplicate characters."

	conversation = add_user_message(conversation, userInput)

	message, err := chat(conversation, system_prompt)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%s\n", message)
	fmt.Println("----------------------------")
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
func chat(conversation []anthropic.MessageParam, system_options ...anthropic.TextBlockParam) (string, error) {

	message_params := anthropic.MessageNewParams{
		MaxTokens: 1024,
		Model:     anthropic.ModelClaude3_5HaikuLatest,
		Messages:  conversation,
	}

	if len(system_options) > 0 {
		message_params.System = system_options // Optional system-level instructions to the model. (e.g. 'You are a helpful assistant.')
	}

	// Make a new Request
	message, err := client.Messages.New(context.TODO(), message_params)
	if err != nil {
		return "", err
	}

	return message.Content[0].Text, nil
}
