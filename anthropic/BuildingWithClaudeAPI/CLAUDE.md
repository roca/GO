# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is a learning repository for the "Building with Claude API" course. It contains multiple independent Go examples demonstrating various Claude API features and patterns, from basic API usage to advanced RAG systems and MCP implementations.

## Project Structure

The repository uses Go workspaces (`go.work`) with each example as an independent module:

- **example1-17**: Progressive examples demonstrating Claude API features
- **go-agent**: Basic agentic workflow with custom tool definitions
- **cli_project**: Full MCP (Model Context Protocol) implementation with client/server architecture
- **genkit-intro**: Firebase Genkit integration example

### Key Example Implementations

- **example11**: Prompt evaluation framework with JSON dataset generation
- **example14**: Tool use patterns with custom time manipulation tools
- **example15**: RAG system with dual search (vector similarity + BM25 lexical search)
- **example16-17**: Extended streaming and advanced API patterns

## Common Development Commands

### Building and Running Examples

Each example is a separate module. Navigate to the example directory and run:

```bash
cd example1  # or any example directory
go run main.go
```

For examples with multiple packages (example14, example15):

```bash
cd example14
go run .
```

### Running the MCP Client/Server

The `cli_project` contains a complete MCP implementation:

```bash
# Terminal 1 - Start MCP server
cd cli_project/server
go run .

# Terminal 2 - Run MCP client
cd cli_project/client
go run .
```

### Docker Operations (MCP Server)

Build and run the MCP server in Docker:

```bash
cd cli_project
docker build -t mcp-server --platform linux/amd64 .
docker run -p 8080:8080 mcp-server
```

### Running Tests

For examples with test files (example13, example14):

```bash
cd example13
go test ./tools -v
```

## Architecture Patterns

### Agent Implementation Pattern

Both `go-agent` and `cli_project/client` follow a similar agentic loop pattern:

```go
type Agent struct {
    client         *anthropic.Client
    getUserMessage func() (string, error)
    tools          []ToolDefinition  // or []*mcp.Tool
}

func (a *Agent) run(ctx context.Context) error {
    conversation := []anthropic.MessageParam{}
    for {
        // 1. Get user input
        // 2. Run inference with Claude API
        // 3. Handle tool calls in loop
        // 4. Display response
    }
}
```

The agent maintains conversation history and handles tool use through multiple API round-trips until Claude returns a final response.

### MCP Architecture

The `cli_project` implements a complete MCP system:

- **Server** (`cli_project/server/`):
  - Exposes tools via `mcp.AddTool()`
  - Serves resources via `AddListDocIDsResource()` and `AddDocContentResource()`
  - Provides prompts via `AddReformatToMarkdownPrompt()`
  - Runs on HTTP with SSE transport at `:8080`

- **Client** (`cli_project/client/`):
  - Connects to MCP server via SSE transport
  - Fetches available tools and resources
  - Interactive CLI with `@` syntax for document references (Ctrl+@ for menu)
  - Integrates MCP tools with Claude API tool calling

### RAG System Architecture (example15)

Document processing pipeline:

1. **Chunking**: `ChunkBySection()` splits on markdown headers (default strategy)
2. **Embedding**: VoyageAI generates vector embeddings (`embedder/embedder.go`)
3. **Indexing**: Dual indexes - `VectorDB` (cosine/euclidean) and `BM25Index` (lexical)
4. **Search**: Query both indexes and merge results with `MergeIndexDBresults()`

Search interface consistency:
```go
type Result struct { Content string; Distance float64 }
Search(query string, topK int, params...) ([]Result, error)
```

### Tool Definition Pattern

Tools follow a consistent schema across examples:

```go
type ToolDefinition struct {
    Name        string
    Description string
    Parameters  map[string]interface{}  // JSON schema
    Execute     func(params map[string]interface{}) (string, error)
}
```

MCP tools use the SDK's `mcp.Tool` type with handlers registered via `mcp.AddTool()`.

## Environment Requirements

### API Keys

Set these environment variables (the SDK looks for them by default):

```bash
export ANTHROPIC_API_KEY="your-key-here"     # Required for all examples
export VOYAGE_API_KEY="your-key-here"        # Required for example15 (RAG)
```

### Dependencies

Each example manages its own dependencies through its `go.mod`. The workspace configuration handles version resolution across all modules.

**Important**: For langchain integrations, this repo uses `github.com/vendasta/langchaingo` (not `tmc/langchaingo`).

## Important Code References

### Streaming Response Pattern

All examples use streaming for real-time responses:

```go
response_stream := client.Messages.NewStreaming(ctx, message_params)
for response_stream.Next() {
    current := response_stream.Current()
    switch current.Type {
    case "content_block_delta":
        // Process text delta
    case "tool_use":
        // Handle tool call
    }
}
```

### MCP Resource Retrieval

Access documents via MCP resources:

```go
transport := &mcp.SSEClientTransport{Endpoint: "http://localhost:8080/"}
session, _ := mcpClient.Connect(ctx, transport, nil)
results, _ := session.ReadResource(ctx, &mcp.ReadResourceParams{
    URI: "docs://documents",
})
```

### Prompt Evaluation Framework (example11)

Three-stage evaluation:
1. Generate test dataset with Claude (JSON array of tasks)
2. Run prompts against each task
3. Grade with dual scoring: syntax validation + model-based evaluation

## Development Notes

- Each example is self-contained and demonstrates specific API features
- The MCP implementation requires both server and client running simultaneously
- RAG examples require the VoyageAI API key for embeddings
- The workspace setup allows running `go mod tidy` at the root to update all modules
- Docker deployment is configured for the MCP server only
