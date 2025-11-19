package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func AddListDocIDsResource(server *mcp.Server) {
	var listDocIdsResource *mcp.Resource = &mcp.Resource{
		Name:        "list_doc",
		Description: "List all available document IDs",
		URI:         "docs://documents",
		MIMEType:    "application/json",
	}

	listDocIdsResourceHandler := func(context.Context, *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
		var ids []*mcp.ResourceContents
		for k, _ := range docs {
			ids = append(ids, &mcp.ResourceContents{
				Text: k,
			})
		}
		return &mcp.ReadResourceResult{
			Contents: ids,
		}, nil
	}
	server.AddResource(listDocIdsResource, listDocIdsResourceHandler)
}

func AddDocContentResource(server *mcp.Server) {
	var docContentResourceTemplate *mcp.ResourceTemplate = &mcp.ResourceTemplate{
		Name:        "fetch_doc",
		Description: "Get the content of a document by its ID",
		URITemplate: "docs://documents/{doc_id}",
		MIMEType:    "text/plain",
	}

	docContentResourceHandler := func(ctx context.Context, request *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {

		pathArgs := strings.Split(request.Params.URI, "/")

		docID := pathArgs[len(pathArgs)-1]

		log.Println("Fetching content for doc ID:", docID)

		content, ok := docs[docID]
		if !ok {
			return nil, fmt.Errorf("Document with ID %s not found", docID)
		}

		return &mcp.ReadResourceResult{
			Contents: []*mcp.ResourceContents{
				{
					Text: content,
				},
			},
		}, nil
	}
	server.AddResourceTemplate(docContentResourceTemplate, docContentResourceHandler)
}
