package main

import "github.com/aws/aws-lambda-go/lambda"

// Event ...
type Event struct {
}

func handler(event Event) (string, error) {
	return "", nil
}

func main() {
	lambda.Start(handler)
}
