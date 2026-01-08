package main

import (
	"context"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/bedrock"
	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigProfile("default"),
		config.WithRegion("us-east-1"),
	)
	client := anthropic.NewClient(
		// bedrock.WithLoadDefaultConfig(context.Background()),
		bedrock.WithConfig(cfg),
	)

	content := "Write me a function that checks a string for duplicate characters in Python."

	println("[user]: " + content)

	message, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		MaxTokens: 1024,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(content)),
		},
		Model:         "global.anthropic.claude-haiku-4-5-20251001-v1:0",
		StopSequences: []string{"```\n"},
		System: []anthropic.TextBlockParam{
			anthropic.TextBlockParam{Text: "You are a Python engineer who writes very concise code.", Type: "text"},
		},
	})

	if err != nil {
		panic(err)
	}

	println("[assistant]: " + message.Content[0].Text + message.StopSequence)
}
