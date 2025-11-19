package main

import (
	"context"
	"log"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

func main() {
	// gemini()
	anthropic()
}

// func gemini() {
// 	ctx := context.Background()
// 	g := genkit.Init(ctx, genkit.WithPlugins(&googlegenai.GoogleAI{}))
//
// 	resp, err := genkit.Generate(ctx, g,
// 		ai.WithPrompt("Why is Genkit awesome?"),
// 		ai.WithModelName("googleai/gemini-2.5-flash"),
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	log.Println(resp.Text())
// }

func anthropic() {
	ctx := context.Background()
	g := genkit.Init(ctx, genkit.WithPlugins(&anthropic.Anthropic{}))

	resp, err := genkit.Generate(ctx, g,
		ai.WithPrompt("Why is Genkit awesome?"),
		ai.WithModelName("anthropic/claude-3-7-sonnet-20250219"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Text())
}
