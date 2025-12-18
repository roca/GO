package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
)

var llm *anthropic.LLM

func init() {
	var err error

	llm, err = anthropic.New(
		anthropic.WithModel("claude-haiku-4-5-20251001"),
	)
	if err != nil {
		log.Fatalf("Failed to create Anthropic LLM: %v", err)
	}
}

func main() {
	prompt := "What is the capital of France?"

	ctx := context.Background()
	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatalf("LLM call failed: %v", err)
	}
	fmt.Printf("Response: %s\n", completion)
}
