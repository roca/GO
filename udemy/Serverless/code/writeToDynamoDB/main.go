package main

import (
	"context"
	"fmt"
	"log"

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
		FileType string `json:"fileType"`
		Images   struct {
			Original Image `json:"original"`
			Resized  Image `json:"resized"`
		} `json:"images"`
	} `json:"results"`
}

func init() {
	sess = session.Must(session.NewSession())
	svc = dynamodb.New(sess)
}

func handler(ctx context.Context, event Event) (string, error) {

	images := event.Results.Images

	res := ""

	// for key, image := range images {
	// 	switch key {
	// 	case "original":
	// 		log.Println("original:", image.String())
	// 		return image.String(), nil
	// 	case "resized":
	// 		log.Println("resized:", image.String())
	// 		return image.String(), nil
	// 	default:

	// 	}
	// }

	log.Println(images)

	err := Put(images.Original, images.Resized)
	if err != nil {
		return res, err
	}

	return "DynamoDB item added", nil
}

// Put into Database
func Put(original Image, resized Image) error {

	_, err := svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("thumbnails"),
		Item: map[string]*dynamodb.AttributeValue{
			"original":  {S: aws.String(original.String())},
			"thumbnail": {S: aws.String(resized.String())},
		},
	})

	return err
}

func main() {
	lambda.Start(handler)
}
