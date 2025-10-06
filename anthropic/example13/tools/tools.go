package tools

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/invopop/jsonschema"
)

type ToolDefinition struct {
	Name        string                         `json:"name"`
	Description string                         `json:"description"`
	InputSchema anthropic.ToolInputSchemaParam `json:"input_schema"`
	Function    func(input json.RawMessage) (string, error)
}

var GetTimeToolDefinition = ToolDefinition{
	Name:        "get_time_tool",
	Description: "Gets then current using a format like '%Y-%M-YD'",
	InputSchema: GetTimeInputSchema,
	Function:    GetTime,
}

type GetTimeInput struct {
	Format string `json:"format" jsonschema_description:"The format to use, using Go time format."`
}

var GetTimeInputSchema = GenerateSchema[GetTimeInput]()

func GenerateSchema[T any]() anthropic.ToolInputSchemaParam {
	var reflector = jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}

	var v T

	schema := reflector.Reflect(v)

	return anthropic.ToolInputSchemaParam{
		Properties: schema.Properties,
	}
}

func GetTime(input json.RawMessage) (string, error) {
	getTimeInput := GetTimeInput{}

	err := json.Unmarshal(input, &getTimeInput)
	if err != nil {
		return "", errors.New("invalid input: " + err.Error())
	}
	format := getValue[string](getTimeInput.Format)
	time := time.Now().Format(format)

	return time, nil
}

func getValue[T any](value any) T {
	v, _ := value.(T)
	return v
}
