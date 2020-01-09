package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"udemy.com/sls/sls-notes-backend/api/models"
)

var sess *session.Session
var svc *dynamodb.DynamoDB
var tableName string

// Event ...
type Event interface{}

func init() {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	svc = dynamodb.New(sess)
	tableName = os.Getenv("NOTES_TABLE")
}

func handler(ctx context.Context, event events.KinesisEvent) error {
	log.Println(event)

	writeRequests :=[]*dynamodb.WriteRequest{}

	for _, record := range event.Records {
		dataText := string(record.Kinesis.Data)

		note := models.Note{}
		if err := json.Unmarshal([]byte(dataText), &note); err != nil {
			return err
		}
		log.Printf("%s Data = %v \n", record.EventName, note)


		writeRequests = append(writeRequests, &dynamodb.WriteRequest{
			PutRequest: &dynamodb.PutRequest{
				Item: map[string]*dynamodb.AttributeValue{
					"user_id":   {S: aws.String(note.UserID)},
					"user_name": {S: aws.String(note.UserName)},
					"note_id":   {S: aws.String(note.NoteID)},
					"timestamp": {N: aws.String(fmt.Sprintf("%d", time.Now().Unix()))},
					"expires":   {N: aws.String(fmt.Sprintf("%d", time.Now().Unix()))},
					"cat":       {S: aws.String(note.Cat)},
					"title":     {S: aws.String(note.Title)},
					"content":   {S: aws.String(note.Content)},
				},
			},
		})
	}
	batch := make(map[string][]*dynamodb.WriteRequest)
	batch[tableName] = writeRequests

	_, err := svc.BatchWriteItem(&dynamodb.BatchWriteItemInput{
		RequestItems: batch,
	})
	if err != nil {
		return nil
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
