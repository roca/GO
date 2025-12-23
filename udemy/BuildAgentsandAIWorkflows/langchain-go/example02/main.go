package main

import (
	"log"

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
	// stringPromptTemplates()
	// standardTemplateDefinition()
	// jinjaPromptTemplates()
	// multilinePromptTemplates()
	// partialVariablePromptTemplates()
	// promptWithModels()
	// chatPromptTemplate()
	usingChatModels()

}
