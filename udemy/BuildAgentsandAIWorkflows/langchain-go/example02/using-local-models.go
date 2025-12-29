package main

import (
	"context"
	"log"

	"github.com/tmc/langchaingo/llms/ollama"
)

func usingLocalModels() {
	ctx := context.Background()

	localModel, err := ollama.New(
		ollama.WithServerURL("https://open-webui-ollama.eksdev.aws.regeneron.com"),
		ollama.WithModel("llama3.2"),
	)
	if err != nil {
		log.Fatal(err)
	}

	prompt := "What is the capital of United States ?"

	response, err := localModel.Call(ctx, prompt)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response from local model: %s", response)
}
