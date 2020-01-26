package main

import (
	"context"
	"log"
	"os"
	"fmt"
	"time"

	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Message struct {
	EventSource string    `json:"eventSource"`
	EventTime   time.Time `json:"eventTime"`
	BucketName  string    `json:"bucketName"`
	ObjectKey   string    `json:"objectKey"`
}


var sess *session.Session
var svc *s3.S3
var bobsBucketName string

func init() {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	svc = s3.New(sess)
	bobsBucketName = os.Getenv("BOBS_BUCKET_NAME")
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
			CopySource: aws.String(fmt.Sprintf("/%s/%s",message.BucketName,message.ObjectKey)),
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

	return message, nil
}

func main() {
	lambda.Start(Handler)
}
