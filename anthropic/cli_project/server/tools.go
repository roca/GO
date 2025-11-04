package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var docs map[string]string = map[string]string{
	"deposition.md":   "This deposition covers the testimony of Angela Smith, P.E.",
	"report.pdf":      "The report details the state of a 20m condenser tower.",
	"financials.docx": "These financials outline the project's budget and expenditures.",
	"outlook.pdf":     "This document presents the projected future performance of the system.",
	"plan.md":         "The plan outlines the steps for the project's implementation.",
	"spec.txt":        "These specifications define the technical requirements for the equipment.",
}

type ReadInput struct {
	DocID string `json:"doc_id" jsonschema:"Id of the document to read"`
}

type EditInput struct {
	DocID  string `json:"doc_id" jsonschema:"Id of the document to read"`
	OldStr string `json:"old_str" jsonschema:"The text to replace. Must match exactly, including whitespace"`
	NewStr string `json:"new_str" jsonschema:"The new text to insert in palce of the old text"`
}

type Output struct {
	Content string `json:"content" jsonschema:"Contents of the document"`
}

func ReadDocument(ctx context.Context, req *mcp.CallToolRequest, input ReadInput) (
	*mcp.CallToolResult,
	Output,
	error,
) {

	content, ok := docs[input.DocID]
	if !ok {
		return nil, Output{}, fmt.Errorf("Doc with id %s not found", input.DocID)
	}

	return nil, Output{Content: content}, nil
}

func EditDocument(ctx context.Context, req *mcp.CallToolRequest, input EditInput) (
	*mcp.CallToolResult,
	Output,
	error,
) {

	content, ok := docs[input.DocID]
	if !ok {
		return nil, Output{}, fmt.Errorf("Doc with id %s not found", input.DocID)
	}

	return nil, Output{Content: strings.ReplaceAll(content, input.OldStr, input.NewStr)}, nil
}
