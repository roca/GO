package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"path/filepath"

	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eventbridge"
)

var sess *session.Session
var svc *eventbridge.EventBridge
var eventBusName string

func init() {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	svc = eventbridge.New(sess)
	eventBusName = os.Getenv("EVENTBUS_NAME")
}

func GetResponseHeaders() map[string]string {

	headers := make(map[string]string)

	headers["Access-Control-Allow-Origin"] = "*"

	return headers

}

func putEvent(message string) error {

	var entries []*eventbridge.PutEventsRequestEntry

	log.Println(message)

	now := time.Now()
	entries = append(entries, &eventbridge.PutEventsRequestEntry{
		Detail:       aws.String(message),
		Source:       aws.String("bob.wakeUp"),
		Time:         &now,
		EventBusName: aws.String(eventBusName),
		DetailType:   aws.String("appRequestSubmitted"),
	})

	output, err := svc.PutEvents(&eventbridge.PutEventsInput{
		Entries: entries,
	})
	if err != nil {
		return err
	}

	log.Println((output))
	return nil

}

type Message struct {
	EventSource string    `json:"eventSource"`
	EventTime   time.Time `json:"eventName"`
	BucketName  string    `json:"bucketName"`
	ObjectKey   string    `json:"objectKey"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, event *events.S3Event) (events.APIGatewayProxyResponse, error) {

	for _, record := range event.Records {
		s3 := record.S3

		message := Message{record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key}
		filename := s3.Object.Key
		suffix := filepath.Ext(filename)

		b, err := json.Marshal(&message)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}

		log.Println(string(b))
		if suffix == ".xlsx" {
			err = putEvent(string(b))
			if err != nil {
				return events.APIGatewayProxyResponse{}, err
			}
		}

		response := events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers:    GetResponseHeaders(),
			Body:       string(b),
		}

		return response, nil
	}

	return events.APIGatewayProxyResponse{}, nil
}

func main() {
	lambda.Start(Handler)
}
