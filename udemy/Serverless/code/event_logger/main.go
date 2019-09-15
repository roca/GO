package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

// Event ...
type Event struct {
	Stage          string
	LambdaFunction string
	LambdaVersion  string
}

func handler(ctx context.Context, event Event) (Event, error) {
	// Comment
	event.LambdaFunction = lambdacontext.FunctionName
	event.LambdaVersion = lambdacontext.FunctionVersion
	return event, nil
}

func main() {
	lambda.Start(handler)
}
