package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eventbridge"
)

var sess *session.Session
var svc *eventbridge.EventBridge

func init() {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	svc = eventbridge.New(sess)
}

func putEvent(jsonData interface{}) error {

	var entries []*eventbridge.PutEventsRequestEntry

	log.Println(jsonData)
	b, err := json.Marshal(&jsonData)
	if err != nil {
		return err
	}

	now := time.Now()
	entries = append(entries, &eventbridge.PutEventsRequestEntry{
		Detail: aws.String(string(b)),
		Source: aws.String("bob.wakeUp"),
		Time:   &now,
	})

	_, err = svc.PutEvents(&eventbridge.PutEventsInput{
		Entries: entries,
	})
	if err != nil {
		return err
	}

	return nil

}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var jsonData interface{}

	if err := json.Unmarshal([]byte(event.Body), &jsonData); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	err := putEvent(jsonData)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{}, nil
}

func main() {
	lambda.Start(Handler)
}
