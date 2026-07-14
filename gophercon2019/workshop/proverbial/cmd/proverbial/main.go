package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jboursiquot/proverbial"
)

func handler(req *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx := context.Background()
	return proverbial.NewHandler().Handle(ctx, req)
}

func main() {
	lambda.Start(handler)
}
