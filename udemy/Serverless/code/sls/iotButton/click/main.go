//assumes you have the following environment variables setup for AWS session creation
// AWS_SDK_LOAD_CONFIG=1
// AWS_ACCESS_KEY_ID=XXXXXXXXXX
// AWS_SECRET_ACCESS_KEY=XXXXXXXX
// AWS_DEFAULT_REGION=us-east-1

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotevents"
	"github.com/aws/aws-sdk-go/service/sns"
)

const (
	DATE_FMT = "02-01-2006"
)

type Device struct {
	Type       string            `json:"type"`
	DeviceID   string            `json:"deviceId"`
	Attributes map[string]string `json:"attributes"`
	DeviceArn  string            `json:"deviceArn"`
}

type StdEvent struct {
	ClickType     string  `json:"clickType"`
	ReportedTime  string  `json:"reportedTime"`
	CertificateID string  `json:"certificateId"`
	RemainingLife float64 `json:"remainingLife"`
	TestMode      bool    `json:"testMode"`
}

type ButtonEvent struct {
	Device   Device   `json:"device"`
	StdEvent StdEvent `json:"stdEvent"`
}

func handler(event *iotevents.Event) (events.APIGatewayProxyResponse, error) {

	log.Printf("Event: %v", event)

	/*
		if err := json.Unmarshal([]byte(event.Body), &note); err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
	*/

	log.Println("creating session")
	sess := session.Must(session.NewSession())
	log.Println("session created")

	svc := sns.New(sess)
	log.Println("service created")

	params := &sns.PublishInput{
		// Hello Sunil. This is a message from Amazon Lambda
		Message:     aws.String("Hello. This is a message from Amazon Lambda"),
		PhoneNumber: aws.String("+12017458446"),
	}
	resp, err := svc.Publish(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		log.Println(err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
		}, err
	}

	// Pretty-print the response data.
	log.Println(resp)

	b, _ := json.Marshal(&resp)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(b),
	}, nil
}

func main() {
	lambda.Start(handler)
}
