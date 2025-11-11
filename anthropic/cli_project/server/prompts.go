package main

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func AddReformatToMarkdownPrompt(server *mcp.Server) {
	var argements []*mcp.PromptArgument = []*mcp.PromptArgument{
		{
			Name:        "doc_id",
			Description: "ID (etc. name) of the document to read",
		},
	}

	prompt := &mcp.Prompt{
		Name:        "reformat_to_markdown_prompt",
		Description: "Get a propmt that asks to reformat a document to markdown format",
		Arguments:   argements,
	}

	promptHandler := func(ctx context.Context, req *mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
		doc_id := req.Params.Arguments["doc_id"]

		tc := &mcp.TextContent{
			Text: fmt.Sprintf("Please reformat the document with ID %s to be in markdown format.", doc_id),
		}

		result := &mcp.GetPromptResult{
			Messages: []*mcp.PromptMessage{
				{
					Role:    "user",
					Content: tc,
				},
			},
		}

		return result, nil
	}

	server.AddPrompt(prompt, promptHandler)
}
