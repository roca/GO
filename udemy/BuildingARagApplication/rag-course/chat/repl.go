package chat

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"rag-course/llm"
	"rag-course/rag"
	"strings"
	"sync"
	"time"
)

type Options struct {
	SystemPromptFile string
}

func RunREPL(ctx context.Context, client *llm.Client, retriever *rag.Retriever, opts Options) error {
	in := bufio.NewScanner(os.Stdin)
	in.Buffer(make([]byte, 0, 64*1024), 1024*1024) // Increase buffer size to 1MB

	history, err := seedHistory(opts.SystemPromptFile)
	if err != nil {
		return err
	}

	fmt.Printf("Chat session started (Model: %s).  Type Q to quit.\n", client.Model())

	for {
		fmt.Print("\n> ")
		if !in.Scan() {
			if err := in.Err(); err != nil {
				return err
			}
			return nil
		}

		input := strings.TrimSpace(in.Text())
		if input == "" {
			continue
		}

		if strings.EqualFold(input, "q") ||
			strings.EqualFold(input, "/exit") ||
			strings.EqualFold(input, "exit") ||
			strings.EqualFold(input, "quit") {
			fmt.Println("Goodbye!")
			return nil
		}

		history = append(history, llm.Message{Role: "user", Content: input})
		spin := startSpinner("Thinking...")
		var stopOnce sync.Once

		turn := history
		if retriever != nil {
			contextText, retErr := retriever.Retrieve(ctx, history)
			if retErr != nil {
				fmt.Fprintln(os.Stderr, "retriever error:", retErr)
			} else if contextText != "" {
				// build a turn with the inline context
				turn = withInlineContext(history, contextText)
			}
		}

		// if len(turn) > 0 {
		// 	fmt.Fprintf(os.Stderr, "\nFinal Prompt:\n\n%s\n\n", turn[len(turn)-1].Content)
		// }

		reply, err := client.ChatStream(ctx, turn, func(s string) {
			stopOnce.Do(spin.Stop)
			fmt.Print(s)
		})

		stopOnce.Do(spin.Stop)
		fmt.Println()

		if err != nil {
			fmt.Println(os.Stderr, "Error:", err)
			history = history[:len(history)-1] // Remove the last user message on error
			continue
		}

		history = append(history, reply)
	}
}

func withInlineContext(history []llm.Message, contextText string) []llm.Message {
	if len(history) == 0 || contextText == "" {
		return history
	}

	last := history[len(history)-1]
	if last.Role != "user" {
		return history
	}

	out := make([]llm.Message, len(history))
	copy(out, history)
	out[len(out)-1] = llm.Message{
		Role:    "user",
		Content: contextText + "\n\n--- Question ---\n\n" + last.Content,
	}
	return out
}

type spinner struct {
	stop chan struct{}
	done chan struct{}
	once sync.Once
}

func startSpinner(label string) *spinner {
	s := &spinner{
		stop: make(chan struct{}),
		done: make(chan struct{}),
	}

	go func() {
		defer close(s.done)
		frames := []string{"|", "/", "-", "\\"}
		t := time.NewTicker(80 * time.Millisecond)
		defer t.Stop()
		i := 0
		for {
			select {
			case <-s.stop:
				fmt.Print("\r\033[2K") // Clear the line
				return
			case <-t.C:
				fmt.Printf("\r%s %s", frames[i%len(frames)], label)
				i++
			}
		}
	}()

	return s
}

func (s *spinner) Stop() {
	s.once.Do(func() {
		close(s.stop)
		<-s.done
	})
}

func seedHistory(systemPromptFile string) ([]llm.Message, error) {
	if systemPromptFile == "" {
		return nil, nil
	}

	data, err := os.ReadFile(systemPromptFile)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("read system prompt file: %w", err)
	}

	content := strings.TrimSpace(string(data))
	if content == "" {
		return nil, nil
	}

	return []llm.Message{
		{
			Role:    "system",
			Content: content,
		},
	}, nil
}
