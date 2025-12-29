package main

import (
	"context"
	"log"

	"github.com/tmc/langchaingo/llms/fake"
)

func UsingFakeLLMs() {

	dummyResponses := []string{
		"Hello!",
		"How are you doing today?",
		"Hope everthing is fine?",
	}

	llm := fake.NewFakeLLM(dummyResponses)
	ctx := context.Background()

	prompt := "Write a poem about fast cars."
	response, err := llm.Call(
		ctx,
		prompt,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response from local model: %s", response)
}
