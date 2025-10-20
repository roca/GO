package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/vendasta/langchaingo/embeddings/voyageai"
)

func main() {
	filePath := "report.md" // Replace with your file path

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	text := string(content)

	// chunks := ChunkByChar(text, 500, 150)
	// chunks := ChunkBySentence(text, 5, 1)
	chunks := ChunkBySectiion(text)

	PrintEmbeddings(chunks)
}

func PrintEmbeddings(texts []string) {
	embedder, err := voyageai.NewVoyageAI(
		voyageai.WithToken(os.Getenv("VOYAGE_API_KEY")),
		voyageai.WithModel("voyage-3-large"),
	)
	if err != nil {
		log.Fatalf("Failed to create VoyageAI embedder: %v", err)
	}

	// Embed the documents
	embeddings, err := embedder.EmbedDocuments(context.Background(), texts)
	if err != nil {
		log.Fatalf("Failed to embed documents: %v", err)
	}

	// Print the embeddings
	for i, embedding := range embeddings {
		fmt.Printf("Embedding for text:\n'%s':\nLength: %d\n-------------------------------\n", texts[i], len(embedding))
	}

}

func ChunkByChar(text string, chunk_size, chunk_over_lap int) []string {
	var chunks []string
	start_idx := 0

	for start_idx < len(text) {
		end_idx := min(start_idx+chunk_size, len(text))

		chunk_text := text[start_idx:end_idx]
		chunks = append(chunks, chunk_text)

		start_idx = len(text)
		if end_idx < len(text) {
			start_idx = end_idx - chunk_over_lap
		}
	}

	return chunks
}

func ChunkBySentence(text string, max_sentences_per_chunk, overlap_sentences int) []string {
	regexPatern := `([.!?])(\s+)`
	re := regexp.MustCompile(regexPatern)
	sentences := re.Split(text, -1)
	var chunks []string

	for start_idx := 0; start_idx < len(sentences); {
		end_idx := min(start_idx+max_sentences_per_chunk, len(sentences))
		current_chunk := sentences[start_idx:end_idx]
		chunks = append(chunks, strings.Join(current_chunk, " "))

		start_idx += max_sentences_per_chunk - overlap_sentences

		if start_idx < 0 {
			start_idx = 0
		}
	}
	return chunks
}

func ChunkBySectiion(text string) []string {
	sections := strings.Split(text, "\n## ")
	return sections
}
