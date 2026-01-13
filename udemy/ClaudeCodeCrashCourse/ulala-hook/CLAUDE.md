# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based project named `ulala-hook`, part of a Udemy course (ClaudeCodeCrashCourse). The project is in early stages with a minimal skeleton structure.

## Technology Stack

- **Language**: Go 1.25.3
- **Module**: `ulala-hook`

## Development Commands

### Building
```bash
go build
```

### Running
```bash
go run main.go
```

### Testing
```bash
go test ./...
```

### Dependencies
```bash
# Add dependencies
go get <package>

# Tidy dependencies
go mod tidy
```

## Architecture

The project currently consists of:
- `main.go`: Entry point with placeholder main function
- `go.mod`: Module definition
- `ulala.wav`: Audio file (likely for audio processing/playback features)

Based on git history, this project has evolved through various stages including:
- AWS Bedrock/Claude API integrations
- Next.js setup
- LangChain-style conversational agents
- Web browser automation

The current minimal state suggests this is being rebuilt or refactored for a new lecture series.
