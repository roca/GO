package ingest

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"rag-course/llm"
	"rag-course/vector"
	"strconv"
	"strings"
	"time"
)

const ImagePathPrefix = "/images/"

var imageExtensions = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
	".webp": true,
	".gif":  true,
}

func IsImage(name string) bool {
	return imageExtensions[strings.ToLower(filepath.Ext(name))]
}

func ProcessImagePath(ctx context.Context, name, description string, opts Options, embedder llm.Embedder, store vector.Store) (int, error) {
	if embedder == nil {
		return 0, errors.New("embedder is required for processing images")
	}

	if store == nil {
		return 0, errors.New("vector store is required for processing images")
	}

	base := filepath.Base(name)
	if !IsImage(base) {
		return 0, errors.New("file is not a supported image type")
	}

	// Create a vector for the image
	desc := strings.TrimSpace(description)
	if desc == "" {
		return 0, errors.New("description is required for image processing")
	}

	size := opts.ChunkSize
	if size <= 0 {
		size = defaultChunkSize
	}

	overlap := opts.ChunkOverlap
	if overlap <= 0 {
		overlap = defaultChunkOverlap
	}

	chunks := chunk(desc, size, overlap)
	if len(chunks) == 0 {
		return 0, errors.New("no chunks created from description")
	}

	vectors, err := embedder.Embed(ctx, chunks)
	if err != nil {
		return 0, fmt.Errorf("failed to embed chunks: %w", err)
	}
	if len(vectors) != len(chunks) {
		return 0, fmt.Errorf("number of vectors (%d) does not match number of chunks (%d)", len(vectors), len(chunks))
	}

	if err := store.DeleteBySource(ctx, base); err != nil {
		return 0, fmt.Errorf("failed to delete existing vectors for source %s: %w", base, err)
	}

	ingestedAt := time.Now().UTC().Format(time.RFC3339)
	imagePath := ImagePathPrefix + base
	docs := make([]vector.Document, len(chunks))
	for i, chunk := range chunks {
		docs[i] = vector.Document{
			ID:      fmt.Sprintf("%s#%d", base, i),
			Content: chunk,
			Metadata: map[string]string{
				"source":      base,
				"type":        "image",
				"image_path":  imagePath,
				"chunk_index": strconv.Itoa(i),
				"chunks":      strconv.Itoa(len(chunks)),
				"ingested_at": ingestedAt,
			},
			Embedding: vectors[i],
		}
	}

	if err := store.Upsert(ctx, docs); err != nil {
		return 0, fmt.Errorf("failed to upsert vectors: %w", err)
	}

	return len(chunks), nil
}
