package main

import (
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var sess *session.Session

func init() {
	sess = session.Must(session.NewSession())
}

func handler(authEvent *events.APIGatewayCustomAuthorizerRequest) (string, error) {

	return "Image successfully resized", nil
}

func main() {
	lambda.Start(handler)

}
