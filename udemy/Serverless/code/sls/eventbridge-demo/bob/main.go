package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"

	svclambda "github.com/aws/aws-sdk-go/service/lambda"
	uuid "github.com/satori/go.uuid"
)

type Message struct {
	EventSource string    `json:"eventSource"`
	EventTime   time.Time `json:"eventTime"`
	BucketName  string    `json:"bucketName"`
	ObjectKey   string    `json:"objectKey"`
}

type FileItem struct {
	FileID          string            `json:"file_id"`
	FileName        string            `json:"file_name"`
	TimeStamp       int64             `json:"timestamp"`
	Expires         int64             `json:"expires"`
	AttributeValues map[string]string `json:"attribute_values"`
}

var sess *session.Session
var svc *s3.S3
var dbSvc *dynamodb.DynamoDB
var lambdaSvc *svclambda.Lambda
var bobsBucketName string
var tableName string

func init() {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	svc = s3.New(sess)
	dbSvc = dynamodb.New(sess)
	lambdaSvc = svclambda.New(sess)

	bobsBucketName = os.Getenv("BOBS_BUCKET_NAME")
	tableName = os.Getenv("FILES_TABLE")
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, event events.CloudWatchEvent) (Message, error) {

	message := Message{}

	if err := json.Unmarshal([]byte(event.Detail), &message); err != nil {
		return Message{}, err
	}

	_, err := svc.CopyObject(
		&s3.CopyObjectInput{
			Bucket:     aws.String(bobsBucketName),
			CopySource: aws.String(fmt.Sprintf("/%s/%s", message.BucketName, message.ObjectKey)),
			Key:        aws.String(message.ObjectKey),
		},
	)
	if err != nil {
		return Message{}, err
	}

	b, err := json.MarshalIndent(&message, "", "\t")
	if err != nil {
		return Message{}, err
	}

	log.Printf(string(b))

	dbErr := processFile(message)
	if dbErr != nil {
		return Message{}, dbErr
	}

	input := &svclambda.InvokeInput{
		FunctionName:   aws.String("parser"),
		InvocationType: aws.String("RequestResponse"),
		Payload:        b,
	}

	_, err = lambdaSvc.Invoke(input)
	if err != nil {
		return Message{}, err
	}

	return message, nil
}

func processFile(message Message) error {

	uuid := uuid.NewV4()

	fileItem := FileItem{}

	fileItem.FileID = fmt.Sprintf("%s:%s", message.ObjectKey, uuid)
	fileItem.FileName = message.ObjectKey
	fileItem.TimeStamp = time.Now().Unix()
	fileItem.Expires = time.Now().AddDate(0, 0, 90).Unix()

	_, err := dbSvc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"file_id":          {S: aws.String(fileItem.FileID)},
			"file_name":        {S: aws.String(fileItem.FileName)},
			"timestamp":        {N: aws.String(fmt.Sprintf("%d", fileItem.TimeStamp))},
			"expires":          {N: aws.String(fmt.Sprintf("%d", fileItem.Expires))},
			"attribute_values": {S: aws.String("{\"field1\":\"value1\"}")},
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
