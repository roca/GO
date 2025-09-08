package main

import (
	"context"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
)

func main() {
	// Install dependencies

	// Load env varaibles
	// SDK looks for 'ANTHROPIC_API_KEY' env by default

	// Create an API Client
	client := anthropic.NewClient()

	// Make a Request
	message, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		MaxTokens: 1024,
		Model:     anthropic.ModelClaude3_5HaikuLatest,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock("What is a quaternion? Answer in one sentence.")),
		},
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%#v\n", message.Content[0].Text)
}
