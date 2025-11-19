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
			Description: "ID (etc. name) of the document to reformeat to markdown",
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
			Text: fmt.Sprintf(`
				Your goal is to reformat a document to be written with mardowdn syntax.

				The id of the document you need to reformat is :
				<document_id>
				%s
				</document_id
 
        Add in headers, bullet points, tables, etc as necessary. Feel free to add in extra text, but don't change the meaning of the report.
        Use the 'edit_document' tool to edit the document. After the document has been edited, respond with the final version of the doc. Don't explain your changes.
				`, doc_id),
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
