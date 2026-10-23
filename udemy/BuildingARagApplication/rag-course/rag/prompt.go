package rag

import (
	"fmt"
	"rag-course/vector"
	"strings"
)

const contextPreamble = `Use the following excerpts from the document collection to answer the question.
	Cite sources by file name when you draw from them. If the excerpts do not address the question, say so
	before answering from general knowledge.`

const unkownSource = "(unknown source)"

func formatContext(hits []vector.Result) string {
	if len(hits) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(contextPreamble)
	sb.WriteString("\n\n--- Excerpts ---\n\n")

	for i, h := range hits {
		source := h.Metadata["source"]
		if source == "" {
			source = unkownSource
		}
		fmt.Fprintf(&sb, "[%d] Source: %s (simularity %.2f)\n%s\n\n",
			i+1, source, h.Score, h.Content)
	}

	return strings.TrimSpace(sb.String())
}
