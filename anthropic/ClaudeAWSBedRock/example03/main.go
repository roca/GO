package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"maps"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/bedrock"
	"github.com/tmc/langchaingo/memory"
)

var llmBedrock *bedrock.LLM

func init() {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigProfile("default"),
		config.WithRegion("us-east-1"),
	)
	client := bedrockruntime.NewFromConfig(cfg) // Assumes AWS credentials are set in the environment

	// Create Bedrock LLM options
	opts := []bedrock.Option{
		bedrock.WithClient(client),
		bedrock.WithModel("global.anthropic.claude-haiku-4-5-20251001-v1:0"),
	}

	// ctx := context.Background()
	llmBedrock, err = bedrock.New(opts...)
	if err != nil {
		log.Fatalf("Failed to create Bedrock LLM: %v", err)
	}

}

func main() {
	// Create conversation buffer memory
	bufferMemory := memory.NewConversationBuffer()
	// Create a conversational chain with the LLM and memory
	conversionalChain := chains.NewConversation(llmBedrock, bufferMemory)

	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🤖 Conversational Chatbot Started")
	fmt.Println("Type 'exit' or 'quit' to end the conversation")

	AddUserMessage(ctx, bufferMemory, "System: You are Python engineer who writes very concise code.")

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

		AddUserMessage(ctx, bufferMemory, userInput)

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
		maps.Copy(inputs, memVars)
		// for k, v := range memVars {
		//	inputs[k] = v
		// }

		// Call the conversational chain
		response, err := conversionalChain.Call(ctx, inputs)

		// Clear thinking message
		fmt.Print("\r🤖 Agent: ")

		if err != nil {
			fmt.Printf("❌️ Error: %v\n", err)
			// Still save the error turn to memory for context
			AddAIMessage(ctx, bufferMemory, fmt.Sprintf("Error: %v", err))
			continue
		}

		output := response["text"].(string)
		fmt.Println(output)

		// Save AI response to memory
		err = AddAIMessage(ctx, bufferMemory, output)
		if err != nil {
			log.Printf("Error saving AI message: %v\n", err)
		}
	}
}

func AddUserMessage(ctx context.Context, bufferMemory *memory.ConversationBuffer, message string) error {
	return bufferMemory.ChatHistory.AddUserMessage(ctx, message)
}

func AddAIMessage(ctx context.Context, bufferMemory *memory.ConversationBuffer, message string) error {
	return bufferMemory.ChatHistory.AddAIMessage(ctx, message)
}
