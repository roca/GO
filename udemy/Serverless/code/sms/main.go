//assumes you have the following environment variables setup for AWS session creation
// AWS_SDK_LOAD_CONFIG=1
// AWS_ACCESS_KEY_ID=XXXXXXXXXX
// AWS_SECRET_ACCESS_KEY=XXXXXXXX
// AWS_DEFAULT_REGION=us-east-1

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type Event struct{}

func handler(ev Event) (events.APIGatewayProxyResponse, error) {

	Info.Println("creating session")
	sess := session.Must(session.NewSession())
	Info.Println("session created")

	svc := sns.New(sess)
	Info.Println("service created")

	params := &sns.PublishInput{
		// Hello Sunil. This is a message from Amazon Lambda
		Message:     aws.String("Hello Sunil. This is a message from Amazon Lambda"),
		PhoneNumber: aws.String("+12017458446"),
	}
	resp, err := svc.Publish(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		Error.Println(err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
		}, err
	}

	// Pretty-print the response data.
	Info.Println(resp)

	b, _ := json.Marshal(&resp)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(b),
	}, nil
}

func main() {
	initLoggers(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	lambda.Start(handler)
}
