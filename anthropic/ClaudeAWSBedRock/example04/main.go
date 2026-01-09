package main

import (
	"context"
	"log"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/bedrock"
	"github.com/aws/aws-sdk-go-v2/config"
)

var client anthropic.Client

func init() {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigProfile("default"),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		log.Fatalf("Error eablishing config: %v", err)
	}

	client = anthropic.NewClient(
		// bedrock.WithLoadDefaultConfig(context.Background()),
		bedrock.WithConfig(cfg),
	)
}

func main() {

	content := "Write me a function that checks a string for duplicate characters in Python."

	println("[user]: " + content)

	messages := []anthropic.MessageParam{
		anthropic.NewUserMessage(anthropic.NewTextBlock(content)),
	}

	message, err := chat(messages)
	if err != nil {
		log.Fatal(err)
	}

	println("[assistant]: " + message.Content[0].Text + message.StopSequence)
}

func chat(messages []anthropic.MessageParam) (*anthropic.Message, error) {

	message, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		MaxTokens:     1024,
		Messages:      messages,
		Model:         "global.anthropic.claude-haiku-4-5-20251001-v1:0",
		StopSequences: []string{"```\n"},
		System: []anthropic.TextBlockParam{
			{Text: "You are a Python engineer who writes very concise code.", Type: "text"},
		},
	})

	if err != nil {
		return nil, err
	}

	return message, nil
}
