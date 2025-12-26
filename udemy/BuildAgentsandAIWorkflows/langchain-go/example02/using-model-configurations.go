package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
)

func usingModelConfigurations() {

	prompt := "Who invented the microphone"

	llm, err := anthropic.New(
		anthropic.WithModel("claude-haiku-4-5-20251001"),
		// anthropic.WithToken(os.Getenv("SOME_API_KEY")), // Default API Key from environment variable ANTHROPIC_API_KEY
		anthropic.With
	)
	if err != nil {
		log.Fatalf("Failed to create Anthropic LLM: %v", err)
	}

	ctx := context.Background()
	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatalf("LLM call failed: %v", err)
	}
	fmt.Printf("Response: %s\n", completion)
}
