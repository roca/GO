package main

import (
	"context"

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

func handler(ctx context.Context) (string, error) {
	return proverbial.NewPicker(ddb).Handle(ctx)
}

func main() {
	lambda.Start(handler)
}
