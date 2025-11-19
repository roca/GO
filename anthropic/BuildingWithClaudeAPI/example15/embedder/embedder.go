package embedder

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/vendasta/langchaingo/embeddings/voyageai"
)

var embedder *voyageai.VoyageAI

func init() {
	var err error
	embedder, err = voyageai.NewVoyageAI(
		voyageai.WithToken(os.Getenv("VOYAGE_API_KEY")),
		voyageai.WithModel("voyage-3-large"),
	)
	if err != nil {
		log.Fatalf("Failed to create VoyageAI embedder: %v", err)
	}
}

func GetEmbeddings(texts []string) ([][]float64, error) {
	// Embed the documents
	embeddings, err := embedder.EmbedDocuments(context.Background(), texts)
	if err != nil {
		return nil, fmt.Errorf("failed to embed documents: %w", err)
	}

	// Convert embeddings to [][]float64

	var result [][]float64
	for _, embedding := range embeddings {
		var embeddingFloats []float64
		for _, value := range embedding {
			embeddingFloats = append(embeddingFloats, float64(value))
		}
		result = append(result, embeddingFloats)
	}

	return result, nil
}

func PrintEmbeddings(texts []string) error {
	// Embed the documents
	embeddings, err := GetEmbeddings(texts)
	if err != nil {
		return fmt.Errorf("failed to embed documents: %w", err)
	}

	// Print the embeddings
	for i, embedding := range embeddings {
		fmt.Printf("Embedding for text %d:\n'%s':\nLength: %d\n-------------------------------\n", i, texts[i], len(embedding))
	}
	return nil
}
