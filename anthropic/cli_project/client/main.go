package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	ctx := context.Background()

	// Create a new client, with no features.:w

	client := mcp.NewClient(&mcp.Implementation{Name: "mcp-client", Version: "v1.0.0"}, nil)

	// Connect to a server over http
	// transport := &mcp.CommandTransport{Command: exec.Command("../server/myserver")}
	transport := &mcp.SSEClientTransport{Endpoint: "http://localhost:8080/"}

	session, err := client.Connect(ctx, transport, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// 	toolResults, err := session.ListTools(ctx, &mcp.ListToolsParams{})
	// 	if err != nil {
	// 		log.Fatalf("ListTools failed: %v", err)
	// 	}
	// 	for _, tool := range toolResults.Tools {
	// 		log.Printf("Tool: %s - %s", tool.Name, tool.Description)
	// 	}

	// Call a tool on the server.
	params := &mcp.CallToolParams{
		Name:      "read_doc_contents",
		Arguments: map[string]any{"doc_id": "plan.md"},
	}
	res, err := session.CallTool(ctx, params)
	if err != nil {
		log.Fatalf("CallTool failed: %v", err)
	}
	if res.IsError {
		log.Fatal("tool failed")
	}
	for _, c := range res.Content {
		log.Print(c.(*mcp.TextContent).Text)
	}
}
