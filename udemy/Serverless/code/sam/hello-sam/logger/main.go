package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event interface{}) (events.APIGatewayProxyResponse, error) {

	b, _ := json.Marshal(event)
	eventJSON := string(b)
	log.Println(eventJSON)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       eventJSON,
	}, nil
}

func main() {
	lambda.Start(handler)
}
