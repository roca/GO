package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
}

func handler(c context.Context, ev Event) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{}, nil
}

func main() {
	lambda.Start(handler)
}
