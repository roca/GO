package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/memory"
)

func conversationalChainsDemo() {

	ctx := context.Background()

	bufferMemory := memory.NewConversationBuffer()

	conversionalChain := chains.NewConversation(llmAnthropic, bufferMemory)

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

		response, err := chains.Run(ctx, conversionalChain, userInput)

		// Clear thinking ...
		fmt.Print("\r🤖 Agent: ")

		if err != nil {
			fmt.Printf("❌️ Error: %v\n", err)
			continue
		}

		fmt.Println(response)
	}
}
