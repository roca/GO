package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(c context.Context) (string, error) {
	return "Hello to " + os.Getenv("MyEnv"), nil
}

func main() {
	lambda.Start(handler)
}
