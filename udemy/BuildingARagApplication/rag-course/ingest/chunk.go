package ingest

import (
	"strings"
	"unicode/utf8"
)

func chunk(text string, size, overlap int) []string {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil
	}

	if len(text) <= size {
		return []string{text}
	}

	if overlap < 0 {
		overlap = 0
	}

	if overlap >= size {
		overlap = size / 2
	}

	threshold := size * 7 / 10

	var chunks []string
	n := len(text)
	start := 0

	for start < n {
		end := start + size
		if end >= n {
			if part := strings.TrimSpace(text[start:]); part != "" {
				chunks = append(chunks, part)
			}
			break
		}

		window := text[start:end]
		switch {
		case strings.LastIndex(window, "\n\n") >= threshold:
			end = start + strings.LastIndex(window, "\n\n") + 2
		case strings.LastIndex(window, ". ") >= threshold:
			end = start + strings.LastIndex(window, ". ") + 2
		case strings.LastIndex(window, " ") >= threshold:
			end = start + strings.LastIndex(window, " ") + 1
		}
		end = nextRuneStart(text, end)
		if part := strings.TrimSpace(text[start:end]); part != "" {
			chunks = append(chunks, part)
		}
		next := end - overlap
		if next <= start {
			next = end
		}
		start = nextRuneStart(text, next)
	}

	return chunks

}

func nextRuneStart(s string, i int) int {
	if i < 0 {
		i = 0
	}
	for i < len(s) && !utf8.RuneStart(s[i]) {
		i++
	}
	return i
}
