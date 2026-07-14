package main

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jboursiquot/proverbial"
)

var sess *session.Session
var ddb *dynamodb.DynamoDB

func init() {
	sess = session.Must(session.NewSession())
	ddb = dynamodb.New(sess)
}

func handler(req *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx := context.Background()
	c := &http.Client{Timeout: time.Second * 10}
	return proverbial.NewChecker(c, ddb).Handle(ctx, req)
}

func main() {
	lambda.Start(handler)
}
