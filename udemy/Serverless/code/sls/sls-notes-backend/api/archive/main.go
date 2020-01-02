package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var sess *session.Session
var svc *dynamodb.DynamoDB
var tableName string

// Event ...
type Event interface{}

func init() {
	sess = session.Must(session.NewSession())
	svc = dynamodb.New(sess)
	tableName = os.Getenv("NOTES_ARCHIVE_TABLE")
}

func handler(ctx context.Context, event events.DynamoDBEvent) {
	for _, record := range event.Records {
		log.Println(record.EventName)
	}
}

func main() {
	lambda.Start(handler)
}
