package main

import (
	"fmt"
	"time"

	"github.com/tmc/langchaingo/prompts"
)

func partialVariablePromptTemplates() {

	templateWithJinja := prompts.PromptTemplate{
		Template:       "I want the {{company_name}} financial report for the date {{date}}.",
		InputVariables: []string{"company_name"},
		TemplateFormat: prompts.TemplateFormatJinja2,
		/* PartialVariables: map[string]any{
			"date": func() string {
				return time.Now().Format("2006-01-02")
			},
		}, */
	}

	// Can also set the variable later!
	templateWithJinja.PartialVariables = map[string]any{
		"date": time.Now().Format("2006-01-02"),
	}

	prompt, _ := templateWithJinja.Format(map[string]any{
		"company_name": "Google",
	})

	fmt.Println(prompt)
}
