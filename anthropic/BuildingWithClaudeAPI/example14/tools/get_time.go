package tools

import (
	"encoding/json"
	"errors"
	"time"
)

var GetTimeToolDefinition = ToolDefinition{
	Name:        "get_time_tool",
	Description: "Gets the current using a format like '%Y-%M-YD'",
	InputSchema: GetTimeInputSchema,
	Function:    GetTime,
}

type GetTimeInput struct {
	Format string `json:"format" jsonschema_description:"The format to use, using Go time format."`
}

var GetTimeInputSchema = GenerateSchema[GetTimeInput]()

func GetTime(input json.RawMessage) (string, error) {
	getTimeInput := GetTimeInput{}

	err := json.Unmarshal(input, &getTimeInput)
	if err != nil {
		return "", errors.New("invalid input: " + err.Error())
	}
	format := getValue[string](getTimeInput.Format)
	current_time := time.Now()

	formatted_time := current_time.Format(format)

	return formatted_time, nil
}

func getValue[T any](value any) T {
	v, _ := value.(T)
	return v
}
