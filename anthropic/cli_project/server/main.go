package main

import (
	"context"
	"log"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Input struct {
	Path string `json:"path" jsonschema:"Path of file to read"`
}

type Output struct {
	Content string `json:"content" jsonschema:"Contents of the file"`
}

func ReadFile(ctx context.Context, req *mcp.CallToolRequest, input Input) (
	*mcp.CallToolResult,
	Output,
	error,
) {

	bytes, err := os.ReadFile(input.Path)
	if err != nil {
		return nil, Output{}, err
	}

	return nil, Output{Content: string(bytes)}, nil
}

func main() {
	// Create a server with a single tool.
	server := mcp.NewServer(&mcp.Implementation{Name: "greeter", Version: "v1.0.0"}, nil)
	mcp.AddTool(server, &mcp.Tool{Name: "read_file", Description: "Reads the content of a file"}, ReadFile)
	// Run the server over stdin/stdout, until the client disconnects.
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
