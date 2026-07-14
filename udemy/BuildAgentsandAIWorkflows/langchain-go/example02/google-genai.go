package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
)

func usingGoogleGenAI() {

	prompt := "Who invented the microphone"

	ctx := context.Background()
	completion, err := llms.GenerateFromSinglePrompt(ctx, llmGenAI, prompt)
	if err != nil {
		log.Fatalf("LLM call failed: %v", err)
	}
	fmt.Printf("Response: %s\n", completion)
}
