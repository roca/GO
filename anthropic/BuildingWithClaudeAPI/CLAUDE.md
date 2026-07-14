# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is a learning repository for the "Building with Claude API" course. It contains multiple independent Go examples demonstrating various Claude API features and patterns, from basic API usage to advanced RAG systems and MCP implementations.

## Project Structure

Go workspace (`go.work`, Go 1.25.3) with each example as an independent module:

- **example1-17**: Progressive examples (basic API calls → tool use → RAG → vision)
- **go-agent**: Basic agentic workflow with custom tool definitions (non-streaming, `client.Messages.New()`)
- **cli_project**: Full MCP (Model Context Protocol) client/server implementation
- **lucidchart-mcp**: Standalone MCP server for Lucidchart API (OAuth 2.0, read-only document tools)
- **genkit-intro**: Firebase Genkit integration example

### Key Examples

- **example11**: Prompt evaluation framework — three-stage pipeline: generate dataset → run prompts → grade with dual scoring (syntax validation + model-based)
- **example13-14**: Tool use patterns with time manipulation tools (have unit tests in `tools/`)
- **example15**: RAG system with dual search (VoyageAI vector embeddings + BM25 lexical search)
- **example16**: Vision capabilities — satellite image analysis for fire risk assessment

## Common Development Commands

```bash
# Run any single-file example
cd example1 && go run main.go

# Run examples with multiple packages
cd example14 && go run .

# MCP: start server (terminal 1), then client (terminal 2)
cd cli_project/server && go run .
cd cli_project/client && go run .

# Tests (example13 and example14 have tests)
cd example13 && go test ./tools -v

# Lucidchart MCP server (requires OAuth setup, runs on :8081)
cd lucidchart-mcp && go run .

# Docker (cli_project MCP server only)
cd cli_project && docker build -t mcp-server --platform linux/amd64 .
docker run -p 8080:8080 mcp-server

# Build and push to ECR (requires AWS_PROFILE=saml)
AWS_PROFILE=saml aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 154919775133.dkr.ecr.us-east-1.amazonaws.com
cd cli_project && docker build -t 154919775133.dkr.ecr.us-east-1.amazonaws.com/app_dev/general:mcp-server --platform linux/amd64 .
docker push 154919775133.dkr.ecr.us-east-1.amazonaws.com/app_dev/general:mcp-server
```

## Architecture Patterns

### Agent Loop Pattern

Both `go-agent` and `cli_project/client` implement the same core loop:

1. Get user input
2. Call Claude API with conversation history + tool definitions
3. If response contains `ToolUseBlock`: execute tool, append result as user message, loop back to step 2
4. If response contains `TextBlock`: display and return to step 1

Conversation history is maintained as `[]anthropic.MessageParam`. Tool results are fed back via `anthropic.NewToolResultBlock()` wrapped in `anthropic.NewUserMessage()`.

**go-agent** uses `client.Messages.New()` (non-streaming) with `ToolDefinition` structs.
**cli_project/client** uses the same pattern but with MCP `*mcp.Tool` types, connecting to the MCP server via SSE to execute tools with `session.CallTool()`.

### Tool Definition Pattern

In `go-agent`, tools use `invopop/jsonschema` for type-safe schema generation:

```go
type ToolDefinition struct {
    Name        string
    Description string
    InputSchema anthropic.ToolInputSchemaParam
    Function    func(input json.RawMessage) (string, error)
}

// Generic schema generator from Go structs with jsonschema_description tags
var schema = GenerateSchema[ReadFileInput]()
```

MCP tools (`cli_project`) use the SDK's `mcp.Tool` type with handlers registered via `mcp.AddTool(server, tool, handler)`.

### MCP Architecture (`cli_project`)

- **Server** (`server/`): HTTP with SSE transport at `:8080`. Exposes tools (read_doc_contents, edit_document, get_jira_ticket, update_jira_ticket, get_confluence_page), resources (`docs://documents`, `docs://documents/{doc_id}`), and prompts.
- **Client** (`client/`): Connects via `mcp.SSEClientTransport`. Interactive CLI with `@` prefix for document references and `Ctrl+@` for document menu. Converts `mcp.Tool` → `anthropic.ToolUnionParam` for Claude API.

**Jira Token Override:** The `get_jira_ticket` and `update_jira_ticket` tools accept an optional `token` parameter that overrides the server's `JIRA_API_TOKEN` env var. This allows per-user token injection at call time.

**Confluence Integration:** The `get_confluence_page` tool fetches a Confluence page by ID (`GET /rest/api/content/{id}?expand=body.storage,version,space`). Accepts an optional `token` parameter that overrides the server's `CONFLUENCE_API_TOKEN` env var.

**Prompts:**
- `set_jira_token` — Takes a `token` argument and returns a message instructing Claude to pass it in the `token` parameter for all subsequent Jira tool calls. Invoke this prompt to configure a user-specific Jira token for the conversation.
- `set_confluence_token` — Takes a `token` argument and returns a message instructing Claude to pass it in the `token` parameter for all subsequent Confluence tool calls.
- `reformat_to_markdown_prompt` — Takes a `doc_id` argument and returns a message instructing Claude to reformat the specified document to markdown using the `edit_document` tool.

### Lucidchart MCP Server (`lucidchart-mcp`)

Standalone MCP server exposing Lucidchart API as read-only tools via OAuth 2.0. Runs on `:8081`.

**Files:**
- `auth.go` — OAuth 2.0 lifecycle: interactive authorization code flow (temp callback server on `:9999`), token refresh (Lucidchart rotates refresh tokens), file-based persistence (`.lucid-tokens.json`, `0600` perms). Thread-safe via `sync.RWMutex` with 5-min refresh buffer.
- `client.go` — `LucidClient` wraps all Lucidchart REST API calls. Central `doRequest()` adds `Authorization: Bearer` and `Lucid-Api-Version: 1` headers automatically, calls `oauth.GetValidToken()` to auto-refresh expired tokens.
- `tools.go` — Three MCP tool handlers following `ToolHandlerFor[In, Out]` pattern:
  - `list_documents` — `POST /documents/search` with keywords/pagination
  - `get_document` — `GET /documents/{id}` returns metadata
  - `export_document` — `GET /documents/{id}` with `Accept` header; PNG returns `mcp.ImageContent`, PDF returns base64-encoded text
- `main.go` — Reads OAuth config from env vars, loads/initiates token flow, registers tools with `ReadOnlyHint: true`, serves SSE on `:8081`

**OAuth first-run flow:** Server prints an authorization URL → user opens in browser → Lucid redirects to `localhost:9999/callback` → server exchanges code for tokens → persists to `.lucid-tokens.json`. Subsequent runs load tokens from file and auto-refresh.

### RAG System (example15)

Pipeline: Chunk (`ChunkBySection` splits on `\n## `) → Embed (VoyageAI `voyage-3-large`) → Index (dual: `VectorDB` for cosine/euclidean + `BM25Index` for lexical) → Search both and merge with `MergeIndexDBresults()`.

### Streaming Pattern

Most examples use streaming responses:

```go
response_stream := client.Messages.NewStreaming(ctx, message_params)
for response_stream.Next() {
    current := response_stream.Current()
    // Handle content_block_delta for text, tool_use for tool calls
}
```

## Environment Requirements

```bash
export ANTHROPIC_API_KEY="..."     # Required for all examples
export VOYAGE_API_KEY="..."        # Required for example15 (RAG embeddings)
export JIRA_API_TOKEN="..."        # Default for cli_project Jira tools (can be overridden per-call via token parameter)
export CONFLUENCE_API_TOKEN="..."  # Default for cli_project Confluence tools (can be overridden per-call via token parameter)
export LUCID_CLIENT_ID="..."       # Required for lucidchart-mcp (OAuth 2.0)
export LUCID_CLIENT_SECRET="..."   # Required for lucidchart-mcp (OAuth 2.0)
```

Optional `lucidchart-mcp` env vars: `LUCID_REDIRECT_URI` (default `http://localhost:9999/callback`), `LUCID_TOKEN_FILE` (default `.lucid-tokens.json`), `LUCID_MCP_ADDR` (default `:8081`).

**Important**: For langchain integrations, this repo uses `github.com/vendasta/langchaingo` (not `tmc/langchaingo`).

## Development Notes

- Each example is self-contained with its own `go.mod`
- The MCP implementation requires both server and client running simultaneously
- The workspace setup allows running `go mod tidy` at the root to update all modules
- Docker deployment is configured for the MCP server only
- ECR repository: `154919775133.dkr.ecr.us-east-1.amazonaws.com/app_dev/general:mcp-server` (us-east-1, `AWS_PROFILE=saml`)
- The Dockerfile is gitignored (contains org-specific SSL/SSH setup)
- The `lucidchart-mcp` server requires a Lucid Enterprise plan for API access; OAuth credentials come from registering an app at [developer.lucid.co](https://developer.lucid.co)
