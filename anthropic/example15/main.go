package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
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
