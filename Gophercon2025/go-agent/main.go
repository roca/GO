package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
)

func main() {
	client := anthropic.NewClient()

	scanner := bufio.NewReader(os.Stdin)

	getUserMessage := func() (string, error) {
		text, _, err := scanner.ReadLine()

		if err != nil {
			return "", err
		}

		return string(text), nil
	}

	tools := []ToolDefinition{ReadFileDefinition}

	agent := NewAgent(&client, getUserMessage, tools)

	err := agent.run(context.TODO())

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
