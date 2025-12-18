# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a learning repository for the Udemy course "Build Agents and AI Workflows with LangChain and Golang". The repository contains Go-based examples using the LangChain Go library (`github.com/tmc/langchaingo`) to build AI applications.

Reference course materials: https://github.com/coderonfleek/langchaingo-course

## Architecture

The repository uses Go workspaces to manage multiple modules. The `go.work` file at the root defines workspace modules, currently including `./langchain-go`.

Each module is self-contained with its own `go.mod` file and can be developed independently while sharing dependencies through the workspace.

### Current Modules

- **langchain-go**: Example application demonstrating basic LangChain Go usage with Anthropic's Claude API

## Environment Setup

API keys are stored in `.secret` file at the root (not tracked in git). Source this file before running applications:

```bash
source .secret
```

Required environment variables:
- `ANTHROPIC_API_KEY`: Required for Anthropic Claude API calls
- `GEMINI_API_KEY`: For Google Gemini integration (if used)

## Development Commands

### Working with Go Workspace

```bash
# Initialize/update workspace
go work sync

# Add a new module to workspace
go work use ./new-module-name
```

### Building and Running

```bash
# Run a specific module
cd langchain-go && go run main.go

# Build a module
cd langchain-go && go build

# Run with environment variables
source .secret && cd langchain-go && go run main.go
```

### Dependency Management

```bash
# Add dependencies to a module
cd langchain-go && go get github.com/package/name

# Tidy dependencies
cd langchain-go && go mod tidy

# Update workspace dependencies
go work sync
```

### Testing

```bash
# Run tests in a module
cd langchain-go && go test ./...

# Run specific test
cd langchain-go && go test -run TestName ./...

# Run with verbose output
cd langchain-go && go test -v ./...
```

## Code Patterns

### LangChain Go Integration

Applications typically follow this pattern:

1. Initialize LLM client in `init()` function (e.g., `anthropic.New()`)
2. Use `llms.GenerateFromSinglePrompt()` for simple text generation
3. Handle errors from both client initialization and LLM calls
4. Use `context.Background()` for LLM operations

The Anthropic client automatically reads `ANTHROPIC_API_KEY` from environment variables when initialized without explicit configuration.
