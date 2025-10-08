package prompt_evaluator

import (
	"encoding/json"
	"example14/ai"
	"example14/tools"
	"fmt"
	"log"

	"github.com/anthropics/anthropic-sdk-go"
)

type PromptEvaluator struct {
}

// RunGetTimeChat method    runs a chat to get the current time using the provided tools
func (pe *PromptEvaluator) RunGetTimeChat(tools []tools.ToolDefinition) (string, string) {
	prompt := ` What is the exact time, formatted as 2006-01-02 15:04:05 ? `
	fmt.Printf("Prompt: %s\n---------------------------------------------\n", prompt)

	var conversation []anthropic.MessageParam

	conversation = ai.Add_user_message(conversation, prompt)
	// conversation = ai.Add_assistant_message(conversation, "```json")

	var stop_sequences []string
	var text string
	var toolResults []anthropic.ContentBlockParamUnion

	// stop_sequences = append(stop_sequences, "```")
	text, toolUseBlock, err := ai.Chat(conversation, 0.7, stop_sequences, tools)
	if err != nil {
		log.Fatalf("Chat error: %v", err)
	}

	for toolUseBlock.ID != "" {
		result := executeTool(tools, toolUseBlock.ID, toolUseBlock.Name, toolUseBlock.Input)
		toolResults = append(toolResults, result)

		conversation = ai.Add_assistant_with_TextAndToolUse_message(conversation, text, toolUseBlock.Name, toolUseBlock.ID, toolUseBlock.Input)
		conversation = ai.Add_user_with_ToolResult_message(conversation, toolResults)

		text, toolUseBlock, err = ai.Chat(conversation, 0.7, nil, tools)
		if err != nil {
			log.Fatalf("Chat error: %v", err)
		}

		if toolUseBlock.ID == "" {
			break
		}

	}

	return prompt, text
}

// executeTool function    executes a tool based on its name and input, returning the result
func executeTool(tools []tools.ToolDefinition, id string, name string, input json.RawMessage) anthropic.ContentBlockParamUnion {

	var toolDefIndx int
	toolFound := false

	for i, tool := range tools {
		if tool.Name == name {
			toolDefIndx = i
			toolFound = true
			break
		}
	}

	if !toolFound {
		return anthropic.NewToolResultBlock(id, "tool not found", true)
	}

	fmt.Printf("Executing tool {%s}. Execution id: {%s}\n", name, id)

	response, err := tools[toolDefIndx].Function(input)

	if err != nil {
		return anthropic.NewToolResultBlock(id, err.Error(), true)
	}

	return anthropic.NewToolResultBlock(id, response, false)
}
