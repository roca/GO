package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

// Event ...
type Event struct {
}

func handler(ctx context.Context, event Event) (Event, error) {

	time.Sleep(3 * time.Second)

	log.Println("Hello World")
	return event, nil
}

func main() {
	lambda.Start(handler)
}
