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

func handler(ctx context.Context, event Event) (string, error) {
	log.Println(event)
	time.Sleep(3 * time.Second)
	return "Hello World", nil
}

func main() {
	lambda.Start(handler)
}
