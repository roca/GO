package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
)

func responseConfiguration() {

	prompt := "Who invented the microphone"

	ctx := context.Background()
	completion, err := llms.GenerateFromSinglePrompt(
		ctx,
		llmAnthropic,
		prompt,
		llms.WithTemperature(0.8),
		llms.WithMaxTokens(100))
	if err != nil {
		log.Fatalf("LLM call failed: %v", err)
	}
	fmt.Printf("Response: %s\n", completion)
}
