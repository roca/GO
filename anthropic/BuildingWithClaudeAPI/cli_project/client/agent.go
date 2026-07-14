package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Agent struct {
	anthropicClient *anthropic.Client
	mcpClient       *mcp.Client
	getUserMessage  func() (string, error)
	tools           []*mcp.Tool
}

func NewAgent(anthropicClient *anthropic.Client, mcpClient *mcp.Client, getUserMessage func() (string, error), tools []*mcp.Tool) *Agent {
	return &Agent{
		anthropicClient: anthropicClient,
		mcpClient:       mcpClient,
		getUserMessage:  getUserMessage,
		tools:           tools,
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
				// fmt.Printf("Tool use result: %+v\n", result.OfToolResult.Content[0].OfText)
				toolResults = append(toolResults, result)
			}
		}

		if len(toolResults) == 0 {
			readUserInput = true
			continue
		} else {
			// fmt.Println("Tool results")
			// fmt.Printf("\t%+v\n", toolResults)
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
				InputSchema: anthropic.ToolInputSchemaParam{
					Type:       "object",
					Properties: any(tool.InputSchema),
				},
			},
		})
	}

	response, err := a.anthropicClient.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     anthropic.ModelClaude3_5Haiku20241022,
		Messages:  conversation,
		MaxTokens: 1024,
		Tools:     anthropicTools,
	})

	return response, err
}

func (a *Agent) executeTool(id string, name string, input json.RawMessage) anthropic.ContentBlockParamUnion {
	var toolDef *mcp.Tool
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

	ctx := context.Background()

	// Connect to a server over http
	// transport := &mcp.CommandTransport{Command: exec.Command("../server/myserver")}
	transport := &mcp.SSEClientTransport{Endpoint: "http://localhost:8080/"}

	session, err := a.mcpClient.Connect(ctx, transport, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	var arguments map[string]any
	err = json.Unmarshal(input, &arguments)
	if err != nil {
		return anthropic.NewToolResultBlock(id, fmt.Sprintf("failed to unmarshal tool input: %s", err.Error()), true)
	}

	fmt.Printf("Executing tool {%s}. Execution id: {%s}\n", name, id)
	params := &mcp.CallToolParams{
		Name:      toolDef.Name,
		Arguments: arguments,
	}

	response, err := session.CallTool(ctx, params)
	if err != nil {
		return anthropic.NewToolResultBlock(id, err.Error(), true)
	}

	if response.IsError {
		return anthropic.NewToolResultBlock(id, "tool execution failed", true)
	}

	var resultStr string
	for _, c := range response.Content {
		resultStr += c.(*mcp.TextContent).Text
	}

	return anthropic.NewToolResultBlock(id, resultStr, false)
}
