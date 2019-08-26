package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Operation string
	Input     struct {
		Operand1 int
		Operand2 int
	}
}

type Response struct {
	Answer int `json:"answer"`
}

func handler(ev Event) (events.APIGatewayProxyResponse, error) {
	Info.Println("Event", ev)
	var answer int

	switch ev.Operation {
	case "add":
		answer = ev.Input.Operand1 + ev.Input.Operand2
	case "subtract":
		answer = ev.Input.Operand1 - ev.Input.Operand2
	case "multiply":
		answer = ev.Input.Operand1 * ev.Input.Operand2
	case "divide":
		answer = ev.Input.Operand1 / ev.Input.Operand2
	default:
		answer = 0
	}

	response := Response{
		Answer: answer,
	}

	b, _ := json.Marshal(&response)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(b),
	}, nil
}

func main() {
	initLoggers(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	lambda.Start(handler)
}
