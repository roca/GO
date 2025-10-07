package main

import (
	"example13/prompt_evaluator"
	"example13/tools"
	"fmt"
)

func main() {
	var toolDefinitions []tools.ToolDefinition

	tool := tools.GetTimeToolDefinition

	toolDefinitions = append(toolDefinitions, tool)

	evaluator := prompt_evaluator.PromptEvaluator{}

	prompt, text := evaluator.RunGetTimeChat(toolDefinitions)

	fmt.Printf("Prompt: %s\n---------------------------------------------\n", prompt)
	fmt.Printf("Response: %s\n", text)

}
