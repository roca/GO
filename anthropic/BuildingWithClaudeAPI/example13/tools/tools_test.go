package tools

import (
	"encoding/json"
	"testing"
)

func TestGetTime(t *testing.T) {
	// Example usage of the GetTime tool
	input := GetTimeInput{
		Format: "2006-01-02 15:04:05",
	}

	jsonInput, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("Error marshalling input: %v", err)
	}

	tool := GetTimeToolDefinition

	result, err := tool.Function(jsonInput)
	if err != nil {
		t.Fatalf("Error getting time: %v", err)
	}
	t.Logf("Current time: %s", result)
}
