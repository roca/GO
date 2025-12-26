package main

import (
	"fmt"

	"github.com/tmc/langchaingo/prompts"
)

func multilineToTemplates() {

	templateString := `
	Display the following list of {{.profession}}

	{{.names}}

	{{if .display_examples}}
	Here are some examples:
	   {{range .display_examples}}
	     - {{.}}
	   {{end}}
	{{end}}
`

	templateWithMultilines := prompts.PromptTemplate{
		Template:       templateString,
		InputVariables: []string{"profession", "names", "display_examples"},
		TemplateFormat: prompts.TemplateFormatGoTemplate,
	}

	prompt, _ := templateWithMultilines.Format(map[string]any{
		"profession": "Footballers",
		"names":      "[Messi, Ronaldo, Neymar, Mbappe]",
		"display_examples": []string{
			"Kaka - AC Milan",
			"Rooney - Manchester United",
		},
	})

	fmt.Println(prompt)
}
