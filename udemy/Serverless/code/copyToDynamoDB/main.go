package main

import (
	"context"
	"errors"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var sess *session.Session
var svc *dynamodb.DynamoDB

type Input struct {
	Comment string `json:"Comment"`
	Results struct {
		FileType string `json:"fileType"`
	} `json:"results"`
}

type Event []Input

type Response struct {
	Region string `json:"region"`
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

func init() {
	sess = session.Must(session.NewSession())
	svc = dynamodb.New(sess)
}

func handler(ctx context.Context, event Event) (Response, error) {

	res := Response{}

	return res, errors.New("No file copied")
}

func main() {
	lambda.Start(handler)
}
