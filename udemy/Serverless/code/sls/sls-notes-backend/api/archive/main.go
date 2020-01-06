package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"udemy.com/sls/sls-notes-backend/api/models"
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

func UnmarshalStreamImage(attribute map[string]events.DynamoDBAttributeValue, out interface{}) error {

	dbAttrMap := make(map[string]*dynamodb.AttributeValue)

	for k, v := range attribute {

		var dbAttr dynamodb.AttributeValue

		bytes, marshalErr := v.MarshalJSON()
		if marshalErr != nil {
			return marshalErr
		}

		json.Unmarshal(bytes, &dbAttr)
		dbAttrMap[k] = &dbAttr
	}

	return dynamodbattribute.UnmarshalMap(dbAttrMap, out)

}

func handler(ctx context.Context, event events.DynamoDBEvent) error {
	var note models.Note

	for _, record := range event.Records {

		if record.EventName == "REMOVE" {

			log.Printf("Processing request data for event ID %s, type %s.\n", record.EventID, record.EventName)

			err := UnmarshalStreamImage(record.Change.OldImage, &note)
			if err != nil {
				return err
			}

			_, err = svc.PutItem(&dynamodb.PutItemInput{
				TableName: aws.String(tableName),
				Item: map[string]*dynamodb.AttributeValue{
					"user_id":   {S: aws.String(note.UserID)},
					"user_name": {S: aws.String(note.UserName)},
					"note_id":   {S: aws.String(note.NoteID)},
					"timestamp": {N: aws.String(fmt.Sprintf("%d", note.TimeStamp))},
					"expires":   {N: aws.String(fmt.Sprintf("%d", note.Expires))},
					"cat":       {S: aws.String(note.Cat)},
					"title":     {S: aws.String(note.Title)},
					"content":   {S: aws.String(note.Content)},
				},
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
