package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
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

func putEvent(message string) error {

	var entries []*eventbridge.PutEventsRequestEntry

	log.Println(message)

	now := time.Now()
	entries = append(entries, &eventbridge.PutEventsRequestEntry{
		Detail:       aws.String(fmt.Sprintf("{\"message\": \"%s\"}", message)),
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

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	queryParams := make(map[string]string)
	for key, value := range event.QueryStringParameters {
		queryParams[key] = value
	}

	err := putEvent(queryParams["message"])
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	headers := make(map[string]string)

	headers["Access-Control-Allow-Origin"] = "*"

	response := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    headers,
		Body:       string(queryParams["message"]),
	}

	return response, nil
}

func main() {
	lambda.Start(Handler)
}
