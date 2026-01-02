package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/prompts"
)

/*

	-	Synopsis Chain: Take a "topic" and generate a "story".
	-	Character Chain: Take our "story" and extract the main "character".
	-	Backstory Chain: Take our main "character" and build a "backstory".

topic -> SynopsisChain -> story -> CharacterChain -> character -> BackstoryChain -> backstory

*/

func sequentialChainsDemo() {

	ctx := context.Background()

	//Synopsis Chain: Take a "topic" and generate a "story".
	synopsisTemplateString := `
		Create a short story synopsis about {{.input}}
		Synopsis:
	`
	synopsisPrompt := prompts.NewPromptTemplate(
		synopsisTemplateString,
		[]string{"input"},
	)
	synopsisChain := chains.NewLLMChain(llmAnthropic, synopsisPrompt)

	// Character Chain: Take our "story" and extract the main "character".
	characterTemplateString := `
		Based on this story synopsis, identify and describe the main character.
		Synopsis: {{.input}}
		Main Character:
	`
	characterPrompt := prompts.NewPromptTemplate(
		characterTemplateString,
		[]string{"input"},
	)
	characterChain := chains.NewLLMChain(llmAnthropic, characterPrompt)
	_ = characterChain

	// Backstory Chain: Take our main "character" and build a "backstory".
	backstoryTemplateString := `
		Create a detailed backstory for the character.
		Character Description: {{.input}}
		Backstory:
	`
	backstoryPrompt := prompts.NewPromptTemplate(
		backstoryTemplateString,
		[]string{"input"},
	)
	backstoryChain := chains.NewLLMChain(llmAnthropic, backstoryPrompt)
	_ = backstoryChain

	sequentialChain, err := chains.NewSimpleSequentialChain([]chains.Chain{
		synopsisChain,
		characterChain,
		backstoryChain,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Run(): single input and single output
	// Predict(): multiple inputs and single output
	// Call(): multiple inputs and multiple outputs

	result, err := chains.Run(
		ctx,
		sequentialChain,
		"A space explorer discovering a new planet",
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
