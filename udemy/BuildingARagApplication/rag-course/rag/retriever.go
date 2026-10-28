package rag

import (
	"context"
	"fmt"
	"rag-course/llm"
	"rag-course/vector"
)

const defaultTopK = 5

type Options struct {
	TopK     int
	Rewriter *Rewriter
}

type Retriever struct {
	embedder llm.Embedder
	store    vector.Store
	rewriter *Rewriter
	topK     int
}

func New(embedder llm.Embedder, store vector.Store, opts *Options) *Retriever {
	topK := opts.TopK
	if topK <= 0 {
		topK = defaultTopK
	}

	return &Retriever{
		embedder: embedder,
		store:    store,
		rewriter: opts.Rewriter,
		topK:     topK,
	}
}

func (r *Retriever) Retrieve(ctx context.Context, history []llm.Message) (string, error) {
	// build the query
	query := r.buildQuery(ctx, history)
	if query == "" {
		return "", nil
	}

	// get the vectors
	vecs, err := r.embedder.Embed(ctx, []string{query})
	if err != nil {
		return "", fmt.Errorf("embedding query: %w", err)
	}
	if len(vecs) == 0 {
		return "", nil
	}

	// get the hits
	hits, err := r.store.Query(ctx, vecs[0], r.topK)
	if err != nil {
		return "", fmt.Errorf("vector query: %w", err)
	}
	if len(hits) == 0 {
		return "", nil
	}

	// return the formatted context
	return formatContext(hits), nil
}

func (r *Retriever) buildQuery(ctx context.Context, history []llm.Message) string {
	if r.rewriter != nil {
		if q, err := r.rewriter.Rewrite(ctx, history); err == nil && q != "" {
			return q
		}
	}
	return lasteUserMessage(history)
}

func lasteUserMessage(history []llm.Message) string {
	for i := len(history) - 1; i >= 0; i-- {
		if history[i].Role == "user" {
			return history[i].Content
		}
	}
	return ""
}
