package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

type JiraTicketInput struct {
	TicketID string `json:"ticket_id" jsonschema:"ID of the Jira ticket to retrieve"`
	Token    string `json:"token,omitempty" jsonschema:"Optional Bearer token to override JIRA_API_TOKEN env var"`
}

type UpdateJiraTicketInput struct {
	TicketID string `json:"ticket_id" jsonschema:"ID of the Jira ticket to update"`
	Comment  string `json:"comment" jsonschema:"The comment text to add to the ticket"`
	Token    string `json:"token,omitempty" jsonschema:"Optional Bearer token to override JIRA_API_TOKEN env var"`
}

type JiraTicket map[string]any

type ReadInput struct {
	DocID string `json:"doc_id" jsonschema:"ID (etc. name) of the document to read"`
}

type EditInput struct {
	DocID  string `json:"doc_id" jsonschema:"ID (etc. name) of the document to read"`
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

	docs[input.DocID] = strings.ReplaceAll(content, input.OldStr, input.NewStr)

	return nil, Output{Content: docs[input.DocID]}, nil
}

func GetJiraTicket(ctx context.Context, req *mcp.CallToolRequest, input JiraTicketInput) (
	*mcp.CallToolResult,
	JiraTicket,
	error,
) {

	// In a real implementation, you would fetch the ticket from Jira's API.
	// Here, we return a mock ticket for demonstration purposes.

	if input.TicketID == "" {
		return nil, JiraTicket{}, fmt.Errorf("TicketID cannot be empty")
	}
	request, err := http.NewRequest("GET", "https://jira.regeneron.com/rest/api/2/issue/"+input.TicketID, nil)
	if err != nil {
		log.Printf("Failed to create Jira API request: %v", err)
		return nil, JiraTicket{}, fmt.Errorf("Failed to create Jira API request: %v", err)
	}

	token := os.Getenv("JIRA_API_TOKEN")
	if input.Token != "" {
		token = input.Token
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}

	log.Println("Executing:", request.URL)
	response, err := client.Do(request)
	if err != nil || response.StatusCode != 200 {
		log.Printf("Failed to fetch ticket from Jira API: %v", err)
		return nil, JiraTicket{}, fmt.Errorf("Failed to fetch ticket from Jira API: %v", err)
	}
	defer response.Body.Close()

	var ticket JiraTicket

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Failed to read Jira ticket response body: %v", err)
		return nil, JiraTicket{}, fmt.Errorf("Failed to read Jira ticket response body: %v", err)
	}

	err = json.Unmarshal(bytes, &ticket)
	if err != nil {
		log.Printf("Failed to decode Jira ticket JSON: %v", err)
		return nil, JiraTicket{}, fmt.Errorf("Failed to decode Jira ticket JSON: %v", err)
	}

	return nil, ticket, nil
}

func UpdateJiraTicket(ctx context.Context, req *mcp.CallToolRequest, input UpdateJiraTicketInput) (
	*mcp.CallToolResult,
	JiraTicket,
	error,
) {

	// Validate inputs
	if input.TicketID == "" {
		return nil, JiraTicket{}, fmt.Errorf("TicketID cannot be empty")
	}
	if input.Comment == "" {
		return nil, JiraTicket{}, fmt.Errorf("Comment cannot be empty")
	}

	// Prepare the comment payload
	commentPayload := map[string]interface{}{
		"body": input.Comment,
	}
	payloadBytes, err := json.Marshal(commentPayload)
	if err != nil {
		log.Printf("Failed to marshal comment payload: %v", err)
		return nil, JiraTicket{}, fmt.Errorf("Failed to marshal comment payload: %v", err)
	}

	// Create POST request to add comment
	url := "https://jira.regeneron.com/rest/api/2/issue/" + input.TicketID + "/comment"
	request, err := http.NewRequest("POST", url, strings.NewReader(string(payloadBytes)))
	if err != nil {
		log.Printf("Failed to create Jira API request: %v", err)
		return nil, JiraTicket{}, fmt.Errorf("Failed to create Jira API request: %v", err)
	}

	token := os.Getenv("JIRA_API_TOKEN")
	if input.Token != "" {
		token = input.Token
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}

	log.Println("Executing:", request.URL)
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Failed to add comment to Jira ticket: %v", err)
		return nil, JiraTicket{}, fmt.Errorf("Failed to add comment to Jira ticket: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 201 && response.StatusCode != 200 {
		log.Printf("Jira API returned status code: %d", response.StatusCode)
		return nil, JiraTicket{}, fmt.Errorf("Jira API returned status code: %d", response.StatusCode)
	}

	var commentResponse JiraTicket

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Failed to read Jira comment response body: %v", err)
		return nil, JiraTicket{}, fmt.Errorf("Failed to read Jira comment response body: %v", err)
	}

	err = json.Unmarshal(bytes, &commentResponse)
	if err != nil {
		log.Printf("Failed to decode Jira comment response JSON: %v", err)
		return nil, JiraTicket{}, fmt.Errorf("Failed to decode Jira comment response JSON: %v", err)
	}

	return nil, commentResponse, nil
}
