package app

import (
	"context"
	"errors"
	"log"
	"os"
	"rag-course/chat"
	"rag-course/config"
	"rag-course/llm"
	"rag-course/vector"
	"rag-course/vector/pgvector"
)

func Run(ctx context.Context, cfg config.Config) error {
	logger := log.New(os.Stderr, "[rag] ", log.LstdFlags)

	client := llm.New(cfg)

	store, err := openStore(ctx, cfg)
	if err != nil {
		logger.Printf("vector store disabled: %v", err)
	}

	if store != nil {
		defer store.Close()
		logger.Printf("vector store ready")
	}

	return chat.RunREPL(ctx, client, chat.Options{
		SystemPromptFile: cfg.SystemPromptFile,
	})
}

func openStore(ctx context.Context, cfg config.Config) (vector.Store, error) {
	if cfg.DatabaseURL == "" {
		return nil, errors.New("Database connection string URL is blank")
	}

	s, err := pgvector.New(ctx, pgvector.Options{
		DSN:          cfg.DatabaseURL,
		EmbeddingDim: cfg.EmbeddingDim,
	})
	if err != nil {
		return nil, err
	}
	return s, nil
}
