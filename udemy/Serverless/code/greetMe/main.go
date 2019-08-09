package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(req *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	lang, ok := req.QueryStringParameters["lang"]
	if !ok {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "'lang' is missing in query string.",
		}, nil
	}
	return events.APIGatewayProxyResponse{}, nil
}

func main() {
	lambda.Start(handler)
}
