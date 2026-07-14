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

	b, err := json.MarshalIndent(&output, "", "\t")
	if err != nil {
		log.Println(string(b))
	}
	return nil

}

type Message struct {
	EventSource string    `json:"eventSource"`
	EventTime   time.Time `json:"eventTime"`
	BucketName  string    `json:"bucketName"`
	ObjectKey   string    `json:"objectKey"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, event *events.S3Event) (events.APIGatewayProxyResponse, error) {

	fileTypes := make(map[string]string)
	fileTypes[".xls"] = "Microsoft Excel 97-2003 Worksheet"
	fileTypes[".xlt"] = "Microsoft Excel 97-2003 Template"
	fileTypes[".xlsx"] = "Excel WorkBook"
	fileTypes[".xltx"] = "Excel Template"

	for _, record := range event.Records {
		s3 := record.S3

		message := Message{record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key}
		filename := s3.Object.Key
		suffix := filepath.Ext(filename)

		b, err := json.MarshalIndent(&message, "", "\t")
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}

		log.Println(string(b))

		if _, ok := fileTypes[suffix]; ok {
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
