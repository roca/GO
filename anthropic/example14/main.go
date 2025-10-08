package main

import (
	"example14/prompt_evaluator"
	"example14/tools"
	"fmt"
)

func main() {
	var toolDefinitions []tools.ToolDefinition

	getTimeTool := tools.GetTimeToolDefinition
	// addDurationTool := tools.AddDurationToolDefinition

	toolDefinitions = append(toolDefinitions, getTimeTool)

	evaluator := prompt_evaluator.PromptEvaluator{}

	_, text := evaluator.RunGetTimeChat(toolDefinitions)

	fmt.Printf("Response: %s\n", text)

}
