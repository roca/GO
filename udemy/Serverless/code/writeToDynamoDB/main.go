package main

import (
	"context"
	"errors"
	"log"

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
	images map[string]Image `json:"images"`
}

func init() {
	sess = session.Must(session.NewSession())
	svc = dynamodb.New(sess)
}

func handler(ctx context.Context, event Event) (string, error) {

	images := event.images

	res := ""

	for key, image := range images {
		switch key {
		case "original":
			log.Println("original:", image.Bucket)
			break
		case "resized":
			log.Println("resized:", image.Bucket)
			break
		default:
			break

		}
	}

	return res, errors.New("No file recorded to dynamodb")
}

func main() {
	lambda.Start(handler)
}
