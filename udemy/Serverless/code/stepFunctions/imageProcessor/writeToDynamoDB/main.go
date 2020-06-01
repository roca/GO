package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
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

func (image Image) String() string {
	if image.Key == "" {
		return ""
	}
	return fmt.Sprintf("%s|%s|%s", image.Region, image.Bucket, image.Key)
}

type Event struct {
	Records []*events.S3EventRecord `json:"Records"`
	Results struct {
		FileType string             `json:"fileType"`
		Images   []map[string]Image `json:"images"`
	} `json:"results"`
}

func init() {
	sess = session.Must(session.NewSession())
	svc = dynamodb.New(sess)
}

func handler(ctx context.Context, event Event) (string, error) {
	var original, resized Image

	images := event.Results.Images

	res := ""
	for _, imageMap := range images {
		for key, image := range imageMap {
			switch key {
			case "original":
				original = image
			case "resized":
				resized = image
			default:
			}
		}
	}

	log.Println(images)

	err := Put(original, resized)
	if err != nil {
		return res, err
	}

	return "DynamoDB item added", nil
}

// Put into Database
func Put(original Image, resized Image) error {

	_, err := svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("IMAGE_TABLE_NAME")),
		Item: map[string]*dynamodb.AttributeValue{
			"original":  {S: aws.String(original.String())},
			"thumbnail": {S: aws.String(resized.String())},
			"timestamp": {S: aws.String(fmt.Sprintf("%d", time.Now().UnixNano()))},
		},
	})

	return err
}

func main() {
	lambda.Start(handler)
}
