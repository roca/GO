package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/llms/bedrock"
)

var llmAnthropic *anthropic.LLM
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
	browsingAgent()
}
