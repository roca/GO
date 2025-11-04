package main

import (
	"log"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {

	// Create a server with a single tool.
	server := mcp.NewServer(&mcp.Implementation{Name: "greeter", Version: "v1.0.0"}, nil)

	read_doc_tool := &mcp.Tool{
		Name:        "read_doc_contents",
		Description: "Read the contents of a document and return it as a string",
	}

	edit_doc_tool := &mcp.Tool{
		Name:        "edit_document",
		Description: "Edit a document by replacing a string in the documents content with a new string",
	}

	mcp.AddTool(server, read_doc_tool, ReadDocument)
	mcp.AddTool(server, edit_doc_tool, EditDocument)

	addr := ":8080"

	// Run the server over http
	log.Printf("MCP servers serving at %s", addr)

	handler := mcp.NewSSEHandler(func(request *http.Request) *mcp.Server {
		url := request.URL.Path
		log.Printf("Handling request for URL %s\n", url)
		return server
	}, nil)
	log.Fatal(http.ListenAndServe(addr, handler))

}
