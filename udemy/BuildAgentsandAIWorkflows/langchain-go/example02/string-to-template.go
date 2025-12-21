package main

import (
	"fmt"

	"github.com/tmc/langchaingo/prompts"
)

func stringPromptTemplates() string {
	simpleTemplate := prompts.NewPromptTemplate(
		"Write a {{.content_type}} about {{.subject}}",
		[]string{"content_type", "subject"},
	)

	templatInput := map[string]any{
		"content_type": "poem",
		"subject":      "cats",
	}

	simplePrompt, _ := simpleTemplate.Format(templatInput)

	fmt.Println(simplePrompt)

	return simplePrompt

}
