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

	_, text := evaluator.RunGetTimeChat(toolDefinitions)

	fmt.Printf("Response: %s\n", text)

}
