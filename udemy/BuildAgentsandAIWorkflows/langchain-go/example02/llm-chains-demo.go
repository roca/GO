package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/prompts"
)

func llmChainsDemo() {
	templateString := `
		You're a helpfull assitant that	translates text.
		Translate the following text from {{.lang1}} to {{.lang2}}

		Text: {{.text}}
		Translation:
	`

	prompt := prompts.NewPromptTemplate(
		templateString,
		[]string{"lang1", "lang2", "text"},
	)
	chainInputs := map[string]any{
		"lang1": "English",
		"lang2": "French",
		"text":  "Well done.",
	}

	llmChain := chains.NewLLMChain(llmAnthropic, prompt)

	// Run(): single input and single output
	// Predict(): multiple inputs and single output
	// Call(): multiple inputs and multiple outputs

	result, err := chains.Predict(context.Background(), llmChain, chainInputs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
