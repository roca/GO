package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/tools"
)

func reActAgentDemo() {
	calculator := tools.Calculator{}

	agentTools := []tools.Tool{calculator}

	agent := agents.NewOneShotAgent(
		llmAnthropic,
		agentTools,
		agents.WithMaxIterations(10),
	)

	executor := agents.NewExecutor(agent)

	mathProblem := "Whats is 25 times 47 plus the square root of 144?"

	ctx := context.Background()
	result, err := executor.Call(ctx, map[string]any{
		"input": mathProblem,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result["output"])

}
