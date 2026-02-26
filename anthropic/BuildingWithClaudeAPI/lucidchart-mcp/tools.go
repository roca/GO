package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- list_documents ---

type ListDocumentsInput struct {
	Keywords   string `json:"keywords"    jsonschema:"Search keywords to find documents by relevance"`
	PageSize   int    `json:"page_size"   jsonschema:"Number of results per page (default 25)"`
	PageNumber int    `json:"page_number" jsonschema:"Page number for pagination (default 1)"`
}

type ListDocumentsOutput struct {
	Documents  []DocumentSummary `json:"documents"   jsonschema:"List of matching documents"`
	TotalHits  int               `json:"total_hits"  jsonschema:"Total number of matching documents"`
	PageSize   int               `json:"page_size"   jsonschema:"Number of results returned"`
	PageNumber int               `json:"page_number" jsonschema:"Current page number"`
}

func ListDocuments(ctx context.Context, req *mcp.CallToolRequest, input ListDocumentsInput) (*mcp.CallToolResult, ListDocumentsOutput, error) {
	if input.PageSize <= 0 {
		input.PageSize = 25
	}
	if input.PageNumber <= 0 {
		input.PageNumber = 1
	}

	result, err := lucidClient.SearchDocuments(ctx, input.Keywords, input.PageSize, input.PageNumber)
	if err != nil {
		return nil, ListDocumentsOutput{}, fmt.Errorf("failed to search documents: %v", err)
	}

	return nil, ListDocumentsOutput{
		Documents:  result.Documents,
		TotalHits:  result.TotalHits,
		PageSize:   result.PageSize,
		PageNumber: result.PageNumber,
	}, nil
}

// --- get_document ---

type GetDocumentInput struct {
	DocumentID string `json:"document_id" jsonschema:"The unique ID of the Lucidchart document"`
}

type GetDocumentOutput struct {
	DocumentID   string `json:"document_id"   jsonschema:"The document ID"`
	Title        string `json:"title"         jsonschema:"The document title"`
	EditURL      string `json:"edit_url"      jsonschema:"URL to edit the document"`
	ViewURL      string `json:"view_url"      jsonschema:"URL to view the document"`
	Product      string `json:"product"       jsonschema:"Lucid product type"`
	LastModified string `json:"last_modified" jsonschema:"Last modification timestamp"`
	Owner        string `json:"owner"         jsonschema:"Document owner"`
	PageCount    int    `json:"page_count"    jsonschema:"Number of pages in the document"`
}

func GetDocument(ctx context.Context, req *mcp.CallToolRequest, input GetDocumentInput) (*mcp.CallToolResult, GetDocumentOutput, error) {
	if input.DocumentID == "" {
		return nil, GetDocumentOutput{}, fmt.Errorf("document_id cannot be empty")
	}

	doc, err := lucidClient.GetDocument(ctx, input.DocumentID)
	if err != nil {
		return nil, GetDocumentOutput{}, fmt.Errorf("failed to get document: %v", err)
	}

	return nil, GetDocumentOutput{
		DocumentID:   doc.DocumentID,
		Title:        doc.Title,
		EditURL:      doc.EditURL,
		ViewURL:      doc.ViewURL,
		Product:      doc.Product,
		LastModified: doc.LastModified,
		Owner:        doc.Owner,
		PageCount:    doc.PageCount,
	}, nil
}

// --- export_document ---

type ExportDocumentInput struct {
	DocumentID string `json:"document_id" jsonschema:"The unique ID of the Lucidchart document to export"`
	Format     string `json:"format"      jsonschema:"Export format: png or pdf (default png)"`
}

type ExportDocumentOutput struct {
	Format string `json:"format" jsonschema:"The export format used"`
}

func ExportDocument(ctx context.Context, req *mcp.CallToolRequest, input ExportDocumentInput) (*mcp.CallToolResult, ExportDocumentOutput, error) {
	if input.DocumentID == "" {
		return nil, ExportDocumentOutput{}, fmt.Errorf("document_id cannot be empty")
	}
	if input.Format == "" {
		input.Format = "png"
	}
	if input.Format != "png" && input.Format != "pdf" {
		return nil, ExportDocumentOutput{}, fmt.Errorf("format must be 'png' or 'pdf', got '%s'", input.Format)
	}

	data, mimeType, err := lucidClient.ExportDocument(ctx, input.DocumentID, input.Format)
	if err != nil {
		return nil, ExportDocumentOutput{}, fmt.Errorf("failed to export document: %v", err)
	}

	output := ExportDocumentOutput{Format: input.Format}

	// PNG: return as native MCP ImageContent
	if input.Format == "png" {
		result := &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.ImageContent{
					Data:     data,
					MIMEType: mimeType,
				},
			},
		}
		return result, output, nil
	}

	// PDF: return as base64-encoded text (no native PDF content type in MCP)
	encoded := base64.StdEncoding.EncodeToString(data)
	contentJSON, _ := json.Marshal(map[string]string{
		"data":      encoded,
		"mime_type": mimeType,
		"format":    input.Format,
	})
	result := &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: string(contentJSON),
			},
		},
	}
	return result, output, nil
}
