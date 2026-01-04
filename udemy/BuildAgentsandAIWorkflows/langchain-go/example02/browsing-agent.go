package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/tools"
	"github.com/vendasta/langchaingo/tools/serpapi"
)

func browsingAgent() {

	search, err := serpapi.New()
	if err != nil {
		log.Fatal(err)
	}

	agentTools := []tools.Tool{search}

	bufferMemory := memory.NewConversationBuffer()
	agent := agents.NewConversationalAgent(
		llmAnthropic,
		agentTools,
		agents.WithMaxIterations(10),
		agents.WithMemory(bufferMemory),
	)

	executor := agents.NewExecutor(agent,
		agents.WithMemory(bufferMemory),
	)

	ctx := context.Background()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\n💬 You: ")

		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading user input: %v\n", err)
		}

		userInput = strings.TrimSpace(userInput)

		if userInput == "" {
			fmt.Println("⚠️ Please enter a message.")
			continue
		}

		fmt.Print("🤖 Agent: ")
		fmt.Print("thinking...")

		response, err := executor.Call(ctx, map[string]any{
			"input": userInput,
		})

		// Clear thinking ...
		fmt.Print("\r🤖 Agent: ")

		if err != nil {
			fmt.Printf("❌️ Error: %v\n", err)
			continue
		}

		fmt.Println(response["output"])
	}
}
