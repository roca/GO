package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
)

func usingChatModels() {

	prompt := "Who invented the microphone"

	ctx := context.Background()
	completion, err := llms.GenerateFromSinglePrompt(ctx, llmAnthropic, prompt)
	if err != nil {
		log.Fatalf("LLM call failed: %v", err)
	}
	fmt.Printf("Response: %s\n", completion)
}
