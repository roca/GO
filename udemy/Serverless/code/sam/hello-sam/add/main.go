package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Operand1 int `json:"num1"`
	Operand2 int `json:"num2"`
}

type Response struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       struct {
		Answer int `json:"answer"`
	} `json:"body"`
}

func handler(ev Event) (events.APIGatewayProxyResponse, error) {
	var answer int

	answer = ev.Operand1 + ev.Operand2

	res := Response{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}
	res.Body.Answer = answer

	b, _ := json.Marshal(&res)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(b),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}, nil
}

func main() {

	lambda.Start(handler)
}
