package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, event interface{}) {

	log.Printf("Event: %v", event)
}

func main() {
	lambda.Start(Handler)
}
