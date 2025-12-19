package main

import (
	"fmt"

	"github.com/tmc/langchaingo/prompts"
)

func standardTemplateDefinition() {

	templateWithProps := prompts.PromptTemplate{
		Template:       "Research {{.topic}} on '{{.website}}'",
		InputVariables: []string{"topic", "website"},
		TemplateFormat: prompts.TemplateFormatGoTemplate,
	}

	prompt, _ := templateWithProps.Format(map[string]any{
		"topic":   "Cookie recipes",
		"website": "The Food Network",
	})

	fmt.Println(prompt)

}
