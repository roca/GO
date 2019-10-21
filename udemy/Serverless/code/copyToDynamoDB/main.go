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

type Image struct {
	Region string `json:"region"`
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

type Event struct {
	Comment string `json:"Comment"`
	Results struct {
		Images []Image `json:"images"`
	} `json:"results"`
}

type Response struct {
	Region string `json:"region"`
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

func init() {
	sess = session.Must(session.NewSession())
	svc = dynamodb.New(sess)
}

func handler(ctx context.Context, event Event) (string, error) {

	images := event.Results.Images

	res := ""

	for _, image := range images {
		switch image.Key {
		case "original":

		}
	}

	return res, errors.New("No file copied")
}

func main() {
	lambda.Start(handler)
}
