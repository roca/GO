package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := "report.md" // Replace with your file path

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	text := string(content)

	chunks := ChunkByChar(text, 500, 150)

	for _, chunk := range chunks {
		fmt.Print(chunk, "\n---\n")
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
