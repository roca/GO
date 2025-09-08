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

	conversation = add_user_message(conversation, "What is a quaternion? Answer in one sentence.")

	// Make a Request
	message, err := chat(conversation)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%#v\n", message)
	fmt.Println("----------------------------")

	// Adds responese from AI and adds on another request!
	conversation = add_assistant_message(conversation, message)
	conversation = add_user_message(conversation, "Give me another sentence.")
	// conversation = add_user_message([]anthropic.MessageParam{}, "Give me another sentence.")

	// Make a new Request
	message, err = chat(conversation)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%#v\n", message)

}

func add_user_message(messages []anthropic.MessageParam, message string) []anthropic.MessageParam {
	conversation := messages
	conversation = append(conversation, anthropic.NewUserMessage(anthropic.NewTextBlock(message)))
	return conversation
}

func add_assistant_message(messages []anthropic.MessageParam, message string) []anthropic.MessageParam {
	conversation := messages
	conversation = append(conversation, anthropic.NewAssistantMessage(anthropic.NewTextBlock(message)))
	return conversation
}

func chat(conversation []anthropic.MessageParam) (string, error) {
	// Make a new Request
	message, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		MaxTokens: 1024,
		Model:     anthropic.ModelClaude3_5HaikuLatest,
		Messages:  conversation,
	})
	if err != nil {
		return "", err
	}

	return message.Content[0].Text, nil
}
