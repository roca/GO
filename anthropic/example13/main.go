package main

import (
	"example13/prompt_evaluator"
	"example13/tools"
	"log"
)

func main() {
	var toolDefinitions []tools.ToolDefinition

	tool := tools.GetTimeToolDefinition

	toolDefinitions = append(toolDefinitions, tool)

	evaluator := prompt_evaluator.PromptEvaluator{}

	prompt, text := evaluator.RunGetTimeChat(toolDefinitions)

	log.Printf("Prompt: %s\n", prompt)
	log.Printf("Response: %s\n", text)

}
