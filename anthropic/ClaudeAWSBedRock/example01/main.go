package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/llms/bedrock"
)

var llmAnthropic *anthropic.LLM
var llmBedrock *bedrock.LLM

func init() {
	var err error

	llmAnthropic, err = anthropic.New(
		anthropic.WithModel("claude-haiku-4-5-20251001"),
		// anthropic.WithToken(os.Getenv("SOME_API_KEY")), // Default API Key from environment variable ANTHROPIC_API_KEY
	)
	if err != nil {
		log.Fatalf("Failed to create Anthropic LLM: %v", err)
	}

	// Create Bedrock LLM options
	opts := []bedrock.Option{
		bedrock.WithModel("amazon.titan-text-lite-v1"),
	}

	// ctx := context.Background()
	llmBedrock, err = bedrock.New(opts...)
	if err != nil {
		log.Fatalf("Failed to create Bedrock LLM: %v", err)
	}

}

func main() {
	prompt := "What is the capital of France?"

	ctx := context.Background()
	completion, err := llms.GenerateFromSinglePrompt(ctx, llmBedrock, prompt)
	if err != nil {
		log.Fatalf("LLM call failed: %v", err)
	}
	fmt.Printf("Response: %s\n", completion)
}
