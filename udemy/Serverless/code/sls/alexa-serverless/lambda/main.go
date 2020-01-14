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
var intents map[string]func(alexa.Request) alexa.Response

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

func HandleFallbackIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Popular Deals", "Popular deal data here")
}

func HandleStopIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Help", "Help regarding the available commands here")
}

func HandleHelpIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Help", "Help regarding the available commands here")
}

func HandleCancelIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Help", "Help regarding the available commands here")
}

func HandleNavigateHomeIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("About", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func HandleGetNewFactIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("About", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func HandleAnotherFactIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("About", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func HandleRepeatIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("About", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func HandleYesIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("About", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func HandleNoIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("About", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func HandleAboutIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("About", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func IntentDispatcher(request alexa.Request) alexa.Response {

	intents = make(map[string]func(alexa.Request) alexa.Response)
	intentName := request.Body.Intent.Name

	intents[alexa.HelpIntent] = HandleHelpIntent
	intents[alexa.CancelIntent] = HandleCancelIntent
	intents[alexa.StopIntent] = HandleStopIntent

	intents["NavigateHomeIntent"] = HandleNavigateHomeIntent
	intents["FallbackIntent"] = HandleFallbackIntent
	intents["GetNewFactIntent"] = HandleGetNewFactIntent
	intents["AnotherFactIntent"] = HandleAnotherFactIntent
	intents["RepeatIntent"] = HandleRepeatIntent
	intents["YesIntent"] = HandleYesIntent
	intents["NoIntent"] = HandleNoIntent

	if intent, ok := intents[intentName]; ok {
		return intent(request)
	}

	return HandleAboutIntent(request)
}

func Handler(request alexa.Request) (alexa.Response, error) {
	return IntentDispatcher(request), nil
}

func main() {
	lambda.Start(Handler)
}
