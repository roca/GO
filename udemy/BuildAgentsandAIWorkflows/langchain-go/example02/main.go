package main

import (
	"context"
	"log"

	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/llms/googleai"
)

var llm *anthropic.LLM
var llmGenAI *googleai.GoogleAI

func init() {
	var err error

	llm, err = anthropic.New(
		anthropic.WithModel("claude-haiku-4-5-20251001"),
	)
	if err != nil {
		log.Fatalf("Failed to create Anthropic LLM: %v", err)
	}

	ctx := context.Background()
	llmGenAI, err = googleai.New(ctx)
}

func main() {
	// stringPromptTemplates()
	// standardTemplateDefinition()
	// jinjaPromptTemplates()
	// multilinePromptTemplates()
	// partialVariablePromptTemplates()
	// promptWithModels()
	// chatPromptTemplate()
	usingChatModels()
	// usingGoogleGenAI()
}
