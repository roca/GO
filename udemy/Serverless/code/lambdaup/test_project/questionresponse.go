package main

/*
   In this example we will up load the zipped copy
   of the executable into the AWS console interface
*/

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Question string
}

type Response struct {
	Question string
	Answer   string
}

func handler(e Event) (Response, error) {
	return Response{
		Question: e.Question,
		Answer:   "I don't know. " + e.Question,
	}, nil
}

func main() {
	lambda.Start(handler)
}
