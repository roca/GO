package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
)

type Agent struct {
	client         *anthropic.Client
	getUserMessage func() (string, error)
	tools          []ToolDefinition
}

func NewAgent(client *anthropic.Client, getUserMessage func() (string, error), tools []ToolDefinition) *Agent {
	return &Agent{
		client:         client,
		getUserMessage: getUserMessage,
		tools:          tools,
	}
}

func (a *Agent) run(ctx context.Context) error {
	conversation := []anthropic.MessageParam{}

	fmt.Println("Chat with claude (Press Ctrl-C to quit)")

	readUserInput := true

	for {
		fmt.Print("> ")

		if readUserInput {
			userInput, err := a.getUserMessage()

			if err != nil {
				return err
			}

			userMessage := anthropic.NewUserMessage(
				anthropic.NewTextBlock(userInput),
			)

			conversation = append(conversation, userMessage)
		}

		response, err := a.runInference(ctx, conversation)
		if err != nil {
			return err
		}

		conversation = append(conversation, response.ToParam())

		var toolResults []anthropic.ContentBlockParamUnion
		for _, block := range response.Content {
			// ? what?
			switch block := block.AsAny().(type) {
			case anthropic.TextBlock:
				// todo stream output.
				// todo prettier colors.
				fmt.Println("┌" + strings.Repeat("─", 60) + "┐")
				fmt.Printf("%+v\n", response.Content[0].Text)
				fmt.Println("└" + strings.Repeat("─", 60) + "┘")
			case anthropic.ToolUseBlock:
				// ? Note that executeTool definitely returns a response. This is fed back into the LLM in the next message.
				result := a.executeTool(block.ID, block.Name, block.Input)
				fmt.Printf("Tool use result: %+v\n", result.OfToolResult.Content[0].OfText)
				toolResults = append(toolResults, result)
			}
		}

		if len(toolResults) == 0 {
			readUserInput = true
			continue
		} else {
			fmt.Println("Tool results")
			fmt.Printf("\t%+v\n", toolResults)
		}

		readUserInput = false
		// ? Let the LLM know what the results of the tool call were.
		conversation = append(conversation, anthropic.NewUserMessage(toolResults...))
	}
}

func (a *Agent) runInference(ctx context.Context, conversation []anthropic.MessageParam) (*anthropic.Message, error) {
	anthropicTools := []anthropic.ToolUnionParam{}

	// must we do this on every inference call?
	for _, tool := range a.tools {
		anthropicTools = append(anthropicTools, anthropic.ToolUnionParam{
			OfTool: &anthropic.ToolParam{
				Name:        tool.Name,
				Description: anthropic.String(tool.Description),
				InputSchema: tool.InputSchema,
			},
		})
	}

	response, err := a.client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     anthropic.ModelClaude3_5Haiku20241022,
		Messages:  conversation,
		MaxTokens: 1024,
		Tools:     anthropicTools,
	})

	return response, err
}

func (a *Agent) executeTool(id string, name string, input json.RawMessage) anthropic.ContentBlockParamUnion {
	var toolDef ToolDefinition
	toolFound := false

	for _, tool := range a.tools {
		if tool.Name == name {
			toolDef = tool
			toolFound = true
			break
		}
	}

	if !toolFound {
		return anthropic.NewToolResultBlock(id, "tool not found", true)
	}

	fmt.Printf("Executing tool {%s}. Execution id: {%s}\n", name, id)

	response, err := toolDef.Function(input)

	if err != nil {
		return anthropic.NewToolResultBlock(id, err.Error(), true)
	}

	return anthropic.NewToolResultBlock(id, response, false)
}
