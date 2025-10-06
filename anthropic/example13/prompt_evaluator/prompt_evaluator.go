package prompt_evaluator

import (
	"example13/ai"
	"example13/tools"
	"log"

	"github.com/anthropics/anthropic-sdk-go"
)

type PromptEvaluator struct {
}

func (pe *PromptEvaluator) RunGetTimeChat(tools []tools.ToolDefinition) (string, string) {
	prompt := `
What is the current time in the format '2006-01-02 15:04:05'?

`
	var conversation []anthropic.MessageParam

	conversation = ai.Add_user_message(conversation, prompt)
	// conversation = ai.Add_assistant_message(conversation, "```json")

	var stop_sequences []string

	//	stop_sequences = append(stop_sequences, "```")
	text, err := ai.Chat(conversation, 0.7, stop_sequences, tools)
	if err != nil {
		log.Fatalf("Chat error: %v", err)
	}

	return prompt, text
}
