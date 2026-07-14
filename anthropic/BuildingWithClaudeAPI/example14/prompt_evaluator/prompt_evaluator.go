package prompt_evaluator

import (
	"encoding/json"
	"example14/ai"
	"example14/tools"
	"fmt"
	"log"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
)

var repeatedString string = strings.Repeat("-", 80)

type PromptEvaluator struct {
}

// RunGetTimeChat method    runs a chat to get the current time using the provided tools
func (pe *PromptEvaluator) RunGetTimeChat(tools []tools.ToolDefinition) (string, string) {
	prompt := ` What date is it 103 days from today ? 
`
	fmt.Printf("Prompt: %s\n%s\n", prompt, repeatedString)

	var conversation []anthropic.MessageParam

	conversation = ai.Add_user_message(conversation, prompt)
	// conversation = ai.Add_assistant_message(conversation, "```json")

	var stop_sequences []string
	var text string

	// stop_sequences = append(stop_sequences, "```")
	text, toolUseBlocks, err := ai.Chat(conversation, 0.0, stop_sequences, tools)
	if err != nil {
		log.Fatalf("Chat error: %v", err)
	}

	for {

		if len(toolUseBlocks) > 0 {
			// fmt.Printf("Tool use blocks found: %d\n", len(toolUseBlocks))
			var toolResults []anthropic.ContentBlockParamUnion
			for _, toolUseBlock := range toolUseBlocks {
				result := executeTool(tools, toolUseBlock.ID, toolUseBlock.Name, toolUseBlock.Input)
				toolResults = append(toolResults, result)
				conversation = ai.Add_assistant_with_TextAndToolUse_message(conversation, text, toolUseBlock.Name, toolUseBlock.ID, toolUseBlock.Input)
			}

			conversation = ai.Add_user_with_ToolResult_message(conversation, toolResults)
		}

		text, toolUseBlocks, err = ai.Chat(conversation, 0.0, nil, tools)
		if err != nil {
			log.Fatalf("Chat error: %v", err)
		}
		// fmt.Println(text)

		if len(toolUseBlocks) == 0 {
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

	fmt.Printf("Executing tool {%s}. Execution id: {%s}, Input %s\n", name, id, string(input))

	response, err := tools[toolDefIndx].Function(input)

	if err != nil {
		return anthropic.NewToolResultBlock(id, err.Error(), true)
	}

	return anthropic.NewToolResultBlock(id, response, false)
}
