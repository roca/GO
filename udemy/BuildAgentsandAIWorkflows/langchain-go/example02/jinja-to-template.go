package main

import (
	"fmt"

	"github.com/tmc/langchaingo/prompts"
)

func jinjaToTemplates() {

	templateWithJinja := prompts.PromptTemplate{
		Template:       "Translate '{{statment}}' from {{lang1}} to {{lang2}}.",
		InputVariables: []string{"statment", "lang1", "lang2"},
		TemplateFormat: prompts.TemplateFormatJinja2,
	}

	prompt, _ := templateWithJinja.Format(map[string]any{
		"statment": "Welcome to China",
		"lang1":    "English",
		"lang2":    "Chinese",
	})

	fmt.Println(prompt)
}
