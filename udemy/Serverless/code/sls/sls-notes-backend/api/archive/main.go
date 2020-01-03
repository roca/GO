package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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

func handler(ctx context.Context, event events.DynamoDBEvent) error {
	for _, record := range event.Records {

		if record.EventName == "REMOVE" {

			log.Printf("Processing request data for event ID %s, type %s.\n", record.EventID, record.EventName)

			// Print new values for attributes of type String
			for name, value := range record.Change.OldImage {
				fmt.Printf("Attribute name: %s, value: %v\n", name, value)
			}

			av, err := dynamodbattribute.MarshalMap(record.Change.OldImage)
			if err != nil {
				return err
			}

			_, err = svc.PutItem(&dynamodb.PutItemInput{
				TableName: aws.String(tableName),
				Item:      av,
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
