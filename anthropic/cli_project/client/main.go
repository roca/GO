package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

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

	scanner := bufio.NewReader(os.Stdin)

	getUserMessage := func() (string, error) {
		text, _, err := scanner.ReadLine()

		if err != nil {
			return "", err
		}

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
