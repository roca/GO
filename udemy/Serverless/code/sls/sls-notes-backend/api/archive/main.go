package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// Event ...
type Event interface{}

func handler(ctx context.Context, event Event) {
	log.Println(event)
}

func main() {
	lambda.Start(handler)
}
