package rag

import (
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
