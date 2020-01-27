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
)

var sess *session.Session
var svc *dynamodb.DynamoDB
var tableName string

type FileItem struct {
	FileID          string `json:"file_id"`
	FileName        string `json:"file_name"`
	TimeStamp       int64  `json:"timestamp"`
	Expires         int64  `json:"expires"`
	AttributeValues string `json:"attribute_values"`
}

func init() {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	svc = dynamodb.New(sess)
	tableName = os.Getenv("FILES_ARCHIVE_TABLE")
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
	var fileItem FileItem

	for _, record := range event.Records {

		if record.EventName == "REMOVE" {

			log.Printf("Processing request data for event ID %s, type %s.\n", record.EventID, record.EventName)

			err := UnmarshalStreamImage(record.Change.OldImage, &fileItem)
			if err != nil {
				return err
			}

			_, err = svc.PutItem(&dynamodb.PutItemInput{
				TableName: aws.String(tableName),
				Item: map[string]*dynamodb.AttributeValue{
					"file_id":           {S: aws.String(fileItem.FileID)},
					"file_name":         {S: aws.String(fileItem.FileName)},
					"timestamp":         {N: aws.String(fmt.Sprintf("%d", fileItem.TimeStamp))},
					"expires":           {N: aws.String(fmt.Sprintf("%d", fileItem.Expires))},
					"attributes_values": {S: aws.String(fileItem.AttributeValues)},
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
