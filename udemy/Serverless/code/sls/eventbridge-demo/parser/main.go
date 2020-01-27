package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event interface{}) error {
	log.Println(event)
	return nil
}

func main() {
	lambda.Start(handler)
}
