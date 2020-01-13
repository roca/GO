package main

/*
	Route: POST /note
*/

import (
	"os"

	"github.com/sls/alexa-serverless/alexa"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var sess *session.Session
var svc *dynamodb.DynamoDB
var tableName string

func init() {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	svc = dynamodb.New(sess)
	tableName = os.Getenv("NOTES_TABLE")
}

func HandleFrontpageDealIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Frontpage Deals", "Frontpage deal data here")
}

func HandlePopularDealIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Popular Deals", "Popular deal data here")
}

func HandleHelpIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Help", "Help regarding the available commands here")
}

func HandleAboutIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("About", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func IntentDispatcher(request alexa.Request) alexa.Response {
	var response alexa.Response
	switch request.Body.Intent.Name {
	case "FrontpageDealIntent":
		response = HandleFrontpageDealIntent(request)
	case "PopularDealIntent":
		response = HandlePopularDealIntent(request)
	case alexa.HelpIntent:
		response = HandleHelpIntent(request)
	case "AboutIntent":
		response = HandleAboutIntent(request)
	default:
		response = HandleAboutIntent(request)
	}
	return response
}

func Handler(request alexa.Request) (alexa.Response, error) {
	return IntentDispatcher(request), nil
}

func main() {
	lambda.Start(Handler)
}
