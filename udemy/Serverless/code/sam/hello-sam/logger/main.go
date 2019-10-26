package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event interface{}) (events.APIGatewayProxyResponse, error) {

	b, _ := json.Marshal(event)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(b),
	}, nil
}

func main() {
	lambda.Start(handler)
}
