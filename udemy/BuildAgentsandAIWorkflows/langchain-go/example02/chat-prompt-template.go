package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/prompts"
)

func chatPromptTemplate() {
	systemPromptTemplate := prompts.NewSystemMessagePromptTemplate(
		"You're to give your responses in {{.language}}",
		[]string{"language"},
	)

	humanMessageTemplate := prompts.NewHumanMessagePromptTemplate(
		"Translate this: {{.some_phrase}}",
		[]string{"some_phrase"},
	)

	chatPromptTemplate := prompts.NewChatPromptTemplate([]prompts.MessageFormatter{
		systemPromptTemplate,
		humanMessageTemplate,
	},
	)

	formattedChatPrompt, _ := chatPromptTemplate.Format(map[string]any{
		"language":    "French",
		"some_phrase": "Thank You!",
	})

	fmt.Println(formattedChatPrompt)

	ctx := context.Background()
	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, formattedChatPrompt)
	if err != nil {
		log.Fatalf("LLM call failed: %v", err)
	}
	fmt.Printf("Response: %s\n", completion)
}
