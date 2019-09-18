package main

import (
	"context"
	"log"
	"os"

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

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Comment
	event.LambdaFunction = lambdacontext.FunctionName
	event.LambdaVersion = lambdacontext.FunctionVersion

	log.Printf("APP_NAME: %s", os.Getenv("APP_NAME"))
	return event, nil
}

func main() {
	lambda.Start(handler)
}
