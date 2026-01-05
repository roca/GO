package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms"
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

	// Create conversation buffer memory
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

	fmt.Println("🤖 Conversational Browsing Agent Started")
	fmt.Println("💡 I can search the web and remember our conversation")
	fmt.Println("Type 'exit' or 'quit' to end the conversation\n")

	for {
		fmt.Print("\n💬 You: ")

		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading user input: %v\n", err)
			continue
		}

		userInput = strings.TrimSpace(userInput)

		if userInput == "" {
			fmt.Println("⚠️ Please enter a message.")
			continue
		}

		// Check for exit commands
		if strings.ToLower(userInput) == "exit" || strings.ToLower(userInput) == "quit" {
			fmt.Println("👋 Goodbye!")
			break
		}

		// Save user message to memory BEFORE calling the agent
		// This ensures the conversation history is available for the next turn
		err = bufferMemory.ChatHistory.AddUserMessage(ctx, userInput)
		if err != nil {
			log.Printf("Error saving user message: %v\n", err)
		}

		// Load memory variables to see what the agent will receive
		memVars, err := bufferMemory.LoadMemoryVariables(ctx, map[string]any{})
		if err != nil {
			log.Printf("Error loading memory: %v\n", err)
		}

		fmt.Print("🤖 Agent: ")
		fmt.Print("thinking...")

		// Prepare inputs including memory history
		inputs := map[string]any{
			"input": userInput,
		}

		// Add memory variables to inputs
		for k, v := range memVars {
			inputs[k] = v
		}

		response, err := executor.Call(ctx, inputs)

		// Clear thinking message
		fmt.Print("\r🤖 Agent: ")

		if err != nil {
			fmt.Printf("❌️ Error: %v\n", err)
			// Still save the error turn to memory for context
			bufferMemory.ChatHistory.AddAIMessage(ctx, fmt.Sprintf("Error: %v", err))
			continue
		}

		output := response["output"].(string)
		fmt.Println(output)

		// Save AI response to memory
		err = bufferMemory.ChatHistory.AddAIMessage(ctx, output)
		if err != nil {
			log.Printf("Error saving AI message: %v\n", err)
		}

		// Debug: Show memory state (optional - comment out if not needed)
		if false {
			history, _ := bufferMemory.ChatHistory.Messages(ctx)
			fmt.Printf("\n[DEBUG] Memory contains %d messages\n", len(history))
			for i, msg := range history {
				msgType := "Unknown"
				switch msg.GetType() {
				case llms.ChatMessageTypeHuman:
					msgType = "Human"
				case llms.ChatMessageTypeAI:
					msgType = "AI"
				}
				fmt.Printf("[DEBUG] Message %d (%s): %s\n", i+1, msgType, msg.GetContent())
			}

			// Show what memory variables were loaded
			fmt.Printf("\n[DEBUG] Memory variables passed to agent:\n")
			for k, v := range memVars {
				fmt.Printf("[DEBUG]   %s: %v\n", k, v)
			}
		}
	}
}
