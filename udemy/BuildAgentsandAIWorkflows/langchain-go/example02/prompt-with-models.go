package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
)

func promptWithModels() {
	prompt := stringPromptTemplates()

	ctx := context.Background()
	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatalf("LLM call failed: %v", err)
	}
	fmt.Printf("Response: %s\n", completion)
}
