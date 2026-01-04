package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/tools"
)

/* A tool must implement the tools.Tool interface:
type Tool interface {
	Name() string
	Description() string
	Call(ctx context.Context, input string) (string, error)
}
Below is a skeleton implementation of a custom tool.
*/

type wordCounterTool struct{}

func (wc wordCounterTool) Name() string {
	return "WordCounter"
}
func (wc wordCounterTool) Description() string {
	return "Counts the number of words in the input text. Input: a string of text. Output: the number of words as an integer."
}
func (wc wordCounterTool) Call(ctx context.Context, input string) (string, error) {
	count := len(strings.Fields(input))
	return fmt.Sprintf("%d", count), nil
}

func usingCustomTool() {

	calculator := tools.Calculator{}
	wordCounter := wordCounterTool{}

	agentTools := []tools.Tool{calculator, wordCounter}

	agent := agents.NewOneShotAgent(
		llmAnthropic,
		agentTools,
		agents.WithMaxIterations(10),
	)

	executor := agents.NewExecutor(agent)

	mathProblem := `
	Count the words in the phrase 
	'Who let the dogs out' 
	and multiply it by the square root of 4.

	Also tell me what tools you used to get the answer.
	`

	ctx := context.Background()
	result, err := executor.Call(ctx, map[string]any{
		"input": mathProblem,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result["output"])

}
