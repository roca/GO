package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var lucidClient *LucidClient

func main() {
	// 1. Configure OAuth from environment
	oauthConfig := OAuthConfig{
		ClientID:     requireEnv("LUCID_CLIENT_ID"),
		ClientSecret: requireEnv("LUCID_CLIENT_SECRET"),
		RedirectURI:  getEnvOrDefault("LUCID_REDIRECT_URI", "http://localhost:9999/callback"),
		Scopes:       []string{"lucidchart.document.content:readonly", "offline_access"},
		TokenFile:    getEnvOrDefault("LUCID_TOKEN_FILE", ".lucid-tokens.json"),
	}

	// 2. Initialize OAuth client and attempt to load existing tokens
	oauth := NewOAuthClient(oauthConfig)
	if err := oauth.LoadTokens(); err != nil {
		log.Printf("No existing tokens found (%v), starting authorization flow...", err)
		if err := oauth.RunAuthorizationFlow(context.Background()); err != nil {
			log.Fatalf("Authorization failed: %v", err)
		}
	} else {
		log.Println("Loaded existing tokens from", oauthConfig.TokenFile)
	}

	// 3. Initialize Lucid API client
	lucidClient = NewLucidClient(oauth)

	// 4. Create MCP server
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "lucidchart-mcp",
		Version: "v1.0.0",
	}, nil)

	// 5. Register tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_documents",
		Description: "Search and list Lucidchart documents by keywords",
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true},
	}, ListDocuments)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_document",
		Description: "Get metadata for a specific Lucidchart document by its ID",
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true},
	}, GetDocument)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "export_document",
		Description: "Export a Lucidchart document as an image (PNG) or PDF",
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true},
	}, ExportDocument)

	// 6. Set up HTTP handler with SSE transport
	handler := mcp.NewSSEHandler(func(request *http.Request) *mcp.Server {
		log.Printf("Handling request for URL %s\n", request.URL.Path)
		return server
	}, nil)

	addr := getEnvOrDefault("LUCID_MCP_ADDR", ":8081")
	log.Printf("Lucidchart MCP server listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}

func requireEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("Required environment variable %s is not set", key)
	}
	return v
}

func getEnvOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
