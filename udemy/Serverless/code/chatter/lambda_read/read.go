package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
}

type Response struct {
	Job string
	Err string
}

func handler(c context.Context, ev Event) (Response, error) {

}

func main() {
	lambda.Start(handler)
}
