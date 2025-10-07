package prompt_evaluator

import (
	"encoding/json"
	"example13/ai"
	"example13/tools"
	"fmt"
	"log"

	"github.com/anthropics/anthropic-sdk-go"
)

type PromptEvaluator struct {
}

func (pe *PromptEvaluator) RunGetTimeChat(tools []tools.ToolDefinition) (string, string) {
	prompt := `
	What is the exact time, formatted as 2006-01-02 15:04:05 ?

`
	var conversation []anthropic.MessageParam

	conversation = ai.Add_user_message(conversation, prompt)
	// conversation = ai.Add_assistant_message(conversation, "```json")

	var stop_sequences []string

	// stop_sequences = append(stop_sequences, "```")
	text, toolUseBlock, err := ai.Chat(conversation, 0.7, stop_sequences, tools)
	if err != nil {
		log.Fatalf("Chat error: %v", err)
	}

	var toolResults []anthropic.ContentBlockParamUnion
	if toolUseBlock != nil {

		result := executeTool(tools, toolUseBlock.ID, toolUseBlock.Name, toolUseBlock.Input)
		// fmt.Printf("Tool use result: %+v\n", result.OfToolResult.Content[0].OfText)
		toolResults = append(toolResults, result)

		// bytes, _ := json.Marshal(toolUseBlock)
		// fmt.Println(string(bytes))
		conversation = ai.Add_assistant_with_TextAndToolUse_message(conversation, text, toolUseBlock.Name, toolUseBlock.ID, toolUseBlock.Input)
	}
	// fmt.Println("Tool results")
	// fmt.Printf("\t%+v\n", toolResults)

	// ? Let the LLM know what the results of the tool call were.
	conversation = append(conversation, anthropic.NewUserMessage(toolResults...))

	text, toolUseBlock, err = ai.Chat(conversation, 0.7, nil, nil)
	if err != nil {
		log.Fatalf("Chat error: %v", err)
	}

	return prompt, text
}

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
