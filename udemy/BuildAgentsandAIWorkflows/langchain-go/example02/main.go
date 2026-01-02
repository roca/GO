package main

import (
	"context"
	"log"

	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/llms/googleai"
)

var llmAnthropic *anthropic.LLM
var llmGenAI *googleai.GoogleAI

func init() {
	var err error

	llmAnthropic, err = anthropic.New(
		anthropic.WithModel("claude-haiku-4-5-20251001"),
		// anthropic.WithToken(os.Getenv("SOME_API_KEY")), // Default API Key from environment variable ANTHROPIC_API_KEY
	)
	if err != nil {
		log.Fatalf("Failed to create Anthropic LLM: %v", err)
	}

	ctx := context.Background()
	llmGenAI, err = googleai.New(
		ctx,
		// googleai.WithAPIKey(os.Getenv("SOME_API_KEY")), // Default API Key from environment variable GOOGLE_API_KEY

	)

}

func main() {
	// stringToTemplates()
	// standardTemplateDefinition()
	// jinjaToTemplates()
	// multilineToTemplates()
	// partialVariableTemplates()
	// promptWithModels()
	// chatPromptTemplate()
	// usingChatModels()
	// usingGoogleGenAI()
	// usingModelConfigurations()
	// responseConfiguration()
	// usingChatModels()
	// usingLocalModels()
	// UsingFakeLLMs()
	// llmChainsDemo()
	// sequentialChainsDemo()
	conversationalChainsDemo()
}
