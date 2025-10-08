package tools

import (
	"encoding/json"
	"errors"
	"time"
)

var AddDurationTool = ToolDefinition{
	Name:        "add_duration_tool",
	Description: "Adds a duration to a given timestamp and returns the new timestamp.",
}

var AddDurationToolDefinition = ToolDefinition{
	Name:        "get_time_tool",
	Description: "Gets the current using a format like '%Y-%M-YD'",
	InputSchema: AddDurationInputSchema,
	Function:    AddDuration,
}

type AddDurationInput struct {
	Format   string `json:"format" jsonschema_description:"The format to use, using Go time format."`
	Duration int    `json:"duration" jsonschema_description:"An integer to specify the duration to add in days"`
}

var AddDurationInputSchema = GenerateSchema[AddDurationInput]()

func AddDuration(input json.RawMessage) (string, error) {
	addDurationInput := AddDurationInput{}

	err := json.Unmarshal(input, &addDurationInput)
	if err != nil {
		return "", errors.New("invalid input: " + err.Error())
	}
	format := getValue[string](addDurationInput.Format)
	current_time := time.Now()

	days := getValue[int](addDurationInput.Duration)

	new_time := current_time.AddDate(0, 0, days)
	formatted_time := new_time.Format(format)

	return formatted_time, nil
}
