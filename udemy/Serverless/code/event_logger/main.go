package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

// Event ...
type Event struct {
	LambdaFunction string
	LambdaVersion  string
}

func handler(c context.Context, event Event) (Event, error) {

	return Event{LambdaFunction: c.functionName, LambdaVersion: c.functionVersion}, nil
}

func main() {
	lambda.Start(handler)
}
