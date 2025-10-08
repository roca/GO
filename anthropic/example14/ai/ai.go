package ai

import (
	"context"
	"encoding/json"
	"example14/tools"
	"time"

	"github.com/anthropics/anthropic-sdk-go"
)

var client anthropic.Client

type ToolUseBlock struct {
	ID    string          `json:"id"`
	Name  string          `json:"name"`
	Input json.RawMessage `json:"input"`
	Type  string          `json:"type"`
}

func init() {
	// Install dependencies
	// Load env varaibles
	// SDK looks for 'ANTHROPIC_API_KEY' env by default

	// Create an API Client
	client = anthropic.NewClient()
}

// Add_user_message function    adds a user message to the conversation
func Add_user_message(messages []anthropic.MessageParam, message string) []anthropic.MessageParam {
	conversation := messages
	conversation = append(conversation, anthropic.NewUserMessage(anthropic.NewTextBlock(message)))
	return conversation
}

// Add_user_with_ToolResultt_message function    adds a user message with tool results to the conversation
func Add_user_with_ToolResult_message(messages []anthropic.MessageParam, toolResults []anthropic.ContentBlockParamUnion) []anthropic.MessageParam {
	conversation := messages
	conversation = append(conversation, anthropic.NewUserMessage(toolResults...))
	return conversation
}

// add_assistant_message function    adds an assistant message to the conversation
func Add_assistant_message(messages []anthropic.MessageParam, message string) []anthropic.MessageParam {
	conversation := messages
	conversation = append(conversation, anthropic.NewAssistantMessage(anthropic.NewTextBlock(message)))
	return conversation
}

// Add_assistant_with_TextAndToolUse_message function    adds an assistant message with a tool use to the conversation
func Add_assistant_with_TextAndToolUse_message(messages []anthropic.MessageParam, message string, tool_name string, tool_id string, tool_input json.RawMessage) []anthropic.MessageParam {
	conversation := messages
	conversation = append(
		conversation,
		anthropic.NewAssistantMessage(anthropic.NewTextBlock(message)),
		anthropic.NewAssistantMessage(anthropic.NewToolUseBlock(tool_id, tool_input, tool_name)),
	)
	return conversation
}

// chat function    sends the conversation to the API and returns the assistant's response
func Chat(
	conversation []anthropic.MessageParam,
	temperature float64,
	stop_sequences []string,
	toolDefinitions []tools.ToolDefinition,
	system_options ...anthropic.TextBlockParam,
) (string, ToolUseBlock, error) {

	//	incomming_message := conversation[len(conversation)-1]
	//
	//	fmt.Printf("%s\n", *incomming_message.Content[0].GetText())
	//	fmt.Println("----------------------------")

	var anthropictools []anthropic.ToolUnionParam

	for _, tool := range toolDefinitions {
		anthropictools = append(anthropictools, anthropic.ToolUnionParam{
			OfTool: &anthropic.ToolParam{
				Name:        tool.Name,
				Description: anthropic.String(tool.Description),
				InputSchema: tool.InputSchema,
			},
		})
	}
	message_params := anthropic.MessageNewParams{
		MaxTokens:     1024,
		Model:         anthropic.ModelClaude3_5HaikuLatest,
		Messages:      conversation,
		Temperature:   anthropic.Float(temperature),
		StopSequences: stop_sequences,
		Tools:         anthropictools,
	}

	if len(system_options) > 0 {
		message_params.System = system_options // Optional system-level instructions to the model. (e.g. 'You are a helpful assistant.')
	}

	ch := make(chan string, 1)
	var toolUseBlock ToolUseBlock

	go func() {
		// Make a new Request

		timeout_ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Minute))
		defer func() {
			cancel()
			// fmt.Println("context cancelled")
			close(ch)
		}()

		response, _ := client.Messages.New(timeout_ctx, message_params)
		if response == nil {
			ch <- "No response from AI"
			return
		}
		if len(response.Content) == 0 {
			ch <- "No content from AI"
			return
		}

		for _, block := range response.Content {
			switch block := block.AsAny().(type) {
			case anthropic.ToolUseBlock:
				toolUseBlock.ID = block.ID
				toolUseBlock.Name = block.Name
				toolUseBlock.Input = block.Input
				toolUseBlock.Type = string(block.Type)
			case anthropic.TextBlock:
				ch <- block.Text
			}
		}

	}()

	var text string

	for t := range ch {
		text = text + t
	}

	return text, toolUseBlock, nil
}
