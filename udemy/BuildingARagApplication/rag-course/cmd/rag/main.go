package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"rag-course/app"
	"rag-course/config"
	"syscall"
)

func main() {
	// We need to:
	// - set up the app
	// - set up config
	// - set up an LLM client
	// - set up the Read-Eval-Print loop (REPL)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := app.Run(ctx, config.Load()); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
