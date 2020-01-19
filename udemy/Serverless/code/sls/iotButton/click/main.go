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
	"github.com/aws/aws-sdk-go/service/sns"
)

const (
	DATE_FMT = "02-01-2006"
)

type ButtonClicked struct {
	ClickType    string `json:"clickType"`
	ReportedTime string `json:"reportedTime"`
}

type DeviceEvent struct {
	ButtonClicked ButtonClicked `json:"buttonClicked"`
}

type DeviceInfo struct {
	Attributes    map[string]string `json:"attributes"`
	DeviceID      string            `json:"deviceId"`
	RemainingLife float32           `json:"remainingLife"`
	Type          string            `json:"type"`
}

type PlacementInfo struct {
	Attributes    map[string]string `json:"attributes"`
	Devices       map[string]string `json:"devices"`
	PlacementName string            `json:"placementName"`
	ProjectName   string            `json:"projectName"`
}

type IoTButtonEvent struct {
	DeviceEvent   DeviceEvent   `json:"deviceEvent"`
	DeviceInfo    DeviceInfo    `json:"deviceInfo"`
	PlacementInfo PlacementInfo `json:"placementInfo`
}

/*
map[
	deviceEvent:map[
		buttonClicked:map[
			clickType:SINGLE
			reportedTime:2020-01-19T17:38:45.362Z
		]
	]
	deviceInfo:map[
		attributes:map[
			deviceTemplateName:lambda-trigger
			placementName:AT-Work
			projectName:IOT-1-Click-demo
			projectRegion:us-west-2
		]
		deviceId:G030PM039134U3AT
		remainingLife:97.65 type:button
	]
	placementInfo:map[
		attributes:map[]
		devices:map[
			lambda-trigger:G030PM039134U3AT
		]
		placementName:AT-Work
		projectName:IOT-1-Click-demo
	]
]
*/

func handler(iotButtonEvent IoTButtonEvent) (events.APIGatewayProxyResponse, error) {

	log.Printf("ClickType: %s", iotButtonEvent.DeviceEvent.ButtonClicked.ClickType)
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
