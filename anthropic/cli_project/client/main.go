package main

import (
	"context"
	"fmt"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var anthropicClient anthropic.Client
var mcpClient *mcp.Client

func init() {
	// Install dependencies
	// Load env varaibles
	// SDK looks for 'ANTHROPIC_API_KEY' env by default

	// Create an API Client
	anthropicClient = anthropic.NewClient()

	// Create a new MCP client, with no features.:w

	mcpClient = mcp.NewClient(&mcp.Implementation{Name: "mcp-client", Version: "v1.0.0"}, nil)
}

func main() {

	getUserMessage := func() (string, error) {
		text := ""

		keyboard.Listen(func(key keys.Key) (stop bool, err error) {
			text += ""

			switch key.Code {
			case keys.Enter:
				return true, nil
			case keys.Backspace:
				if len(text) > 0 {
					text = text[:len(text)-1]
					fmt.Print("\b \b")
				}
				return false, nil
			case keys.Space:
				fmt.Print(" ")
				text += " "
				return false, nil
			case keys.RuneKey:
				fmt.Print(key.String())
				text += key.String()
				return false, nil
			default:
				return false, nil
			}

		})

		return string(text), nil
	}

	// tools := []ToolDefinition{ReadFileDefinition}
	tools := []*mcp.Tool{}

	read_doc_tool := &mcp.Tool{
		Name:        "read_doc_contents",
		Description: "Read the contents of a document and return it as a string",
	}

	edit_doc_tool := &mcp.Tool{
		Name:        "edit_document",
		Description: "Edit a document by replacing a string in the documents content with a new string",
	}

	tools = append(tools, read_doc_tool, edit_doc_tool)

	agent := NewAgent(&anthropicClient, mcpClient, getUserMessage, tools)

	err := agent.run(context.TODO())

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
