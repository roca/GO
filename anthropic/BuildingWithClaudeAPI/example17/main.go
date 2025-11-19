package main

import (
	"encoding/base64"
	"example17/ai"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
)

var prompt string = `
Summarize the document in one sentence.
`

func main() {

	byteData, err := ReadFile("earth.pdf")
	if err != nil {
		panic(err)
	}
	encodedString := base64.StdEncoding.EncodeToString(byteData)

	// fmt.Println("Base64 Encoded String:", encodedString)

	var conversation []anthropic.MessageParam

	conversation = ai.Add_user_with_PDF_And_message(conversation, prompt, encodedString)
	text, _, _ := ai.Chat(conversation, 1.0, nil, nil)
	fmt.Println("AI Response:", text)
}

func ReadFile(path string) ([]byte, error) {

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return content, nil
}
