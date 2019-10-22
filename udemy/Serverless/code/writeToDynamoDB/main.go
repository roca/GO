package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
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
	Images map[string]Image `json:"images"`
}

func init() {
	sess = session.Must(session.NewSession())
	svc = dynamodb.New(sess)
}

func handler(ctx context.Context, event Event) (string, error) {

	images := event.Images

	res := ""

	for key, image := range images {
		s3Local := fmt.Sprintf("%s%s%s", image.Region, image.Bucket, image.Key)
		switch key {
		case "original":
			log.Println("original:", s3Local)
			return s3Local, nil
		case "resized":
			log.Println("resized:", s3Local)
			return s3Local, nil
		default:

		}
	}

	return res, errors.New("No file recorded to dynamodb")
}

// Put into Database
func Put(images map[string]Image) error {

	_, err := svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("thumbnails"),
		Item: map[string]*dynamodb.AttributeValue{
			"original":  {S: aws.String(images["original"])},
			"thumbnail": {S: aws.String(images["resized"])},
		},
	})

	return err
}

func main() {
	lambda.Start(handler)
}
