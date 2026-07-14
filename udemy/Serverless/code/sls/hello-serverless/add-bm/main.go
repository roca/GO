package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

func handler(event Event) (events.APIGatewayProxyResponse, error) {

	response := struct {
		Message   string `json:"message"`
		Timestamp int64  `json:"timestamp"`
	}{
		Message:   fmt.Sprintf("%d + %d =  %d", event.Num1, event.Num2, (event.Num1 + event.Num2)),
		Timestamp: time.Now().Unix(),
	}

	b, _ := json.Marshal(&response)

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
