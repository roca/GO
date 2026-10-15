# AGENTS.md

## Project Overview

Go-based CLI chat application using OpenAI's Go SDK. It's a RAG (Retrieval-Augmented Generation) course project with a simple REPL interface.

## Commands

```bash
# Run the application
go run ./cmd/rag

# Build
go build -o rag ./cmd/rag
```

## Environment

The app loads `.env` automatically via `godotenv`. Required variables:

- `OPENAI_API_KEY` - API key (required)
- `OPENAI_BASE_URL` - Override default `https://api.openai.com/v1` (e.g., `http://localhost:11434/v1` for Ollama)
- `OPENAI_MODEL` - Defaults to `gpt-4o-mini`
- `SYSTEM_PROMPT_FILE` - Optional path to system prompt file

## Structure

- `cmd/rag/main.go` - Entry point
- `app/app.go` - Orchestrates config, LLM client, and REPL
- `config/config.go` - Loads env vars with `.env` support
- `llm/client.go` - OpenAI SDK wrapper with streaming
- `chat/repl.go` - CLI chat loop with history

## Notes

- No tests in this repo
- No lint/format config files
- Uses streaming chat responses
- Exit with `q`, `quit`, `exit`, or `/exit`
