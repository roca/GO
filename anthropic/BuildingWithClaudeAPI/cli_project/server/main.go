package main

import (
	"log"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {

	// Create a server with a single tool.
	server := mcp.NewServer(&mcp.Implementation{Name: "greeter", Version: "v1.0.0"}, nil)

	AddListDocIDsResource(server)
	AddDocContentResource(server)
	AddJiraTokenPrompt(server)
	AddConfluenceTokenPrompt(server)
	AddReformatToMarkdownPrompt(server)

	handler := mcp.NewSSEHandler(func(request *http.Request) *mcp.Server {
		url := request.URL.Path
		log.Printf("Handling request for URL %s\n", url)
		return server
	}, nil)

	read_doc_tool := &mcp.Tool{
		Name:        "read_doc_contents",
		Description: "Read the contents of a document and return it as a string",
	}

	edit_doc_tool := &mcp.Tool{
		Name:        "edit_document",
		Description: "Edit a document by replacing a string in the documents content with a new string",
	}

	get_jira_ticket_tool := &mcp.Tool{
		Name:        "get_jira_ticket",
		Description: "Fetch details of a Jira ticket given its TicketID",
	}

	update_jira_ticket_tool := &mcp.Tool{
		Name:        "update_jira_ticket",
		Description: "Update a Jira ticket by adding a comment to it",
	}

	get_confluence_page_tool := &mcp.Tool{
		Name:        "get_confluence_page",
		Description: "Fetch a Confluence page by its page ID, returning the page content, version, and space info",
	}

	mcp.AddTool(server, read_doc_tool, ReadDocument)
	mcp.AddTool(server, edit_doc_tool, EditDocument)
	mcp.AddTool(server, get_jira_ticket_tool, GetJiraTicket)
	mcp.AddTool(server, update_jira_ticket_tool, UpdateJiraTicket)
	mcp.AddTool(server, get_confluence_page_tool, GetConfluencePage)

	addr := ":8080"

	// Run the server over http
	log.Printf("MCP servers serving at %s", addr)

	log.Fatal(http.ListenAndServe(addr, handler))

}
