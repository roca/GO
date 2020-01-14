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


func IntentDispatcher(request alexa.Request) alexa.Response {

	intents = make(map[string]func(alexa.Request) alexa.Response)
	intentName := request.Body.Intent.Name

	intents[alexa.HelpIntent] = alexa.HandleHelpIntent
	intents[alexa.CancelIntent] = alexa.HandleCancelIntent
	intents[alexa.StopIntent] = alexa.HandleStopIntent

	intents["NavigateHomeIntent"] = alexa.HandleNavigateHomeIntent
	intents["FallbackIntent"] = alexa.HandleFallbackIntent
	intents["GetNewFactIntent"] = alexa.HandleGetNewFactIntent
	intents["AnotherFactIntent"] = alexa.HandleAnotherFactIntent
	intents["RepeatIntent"] = alexa.HandleRepeatIntent
	intents["YesIntent"] = alexa.HandleYesIntent
	intents["NoIntent"] = alexa.HandleNoIntent

	if intent, ok := intents[intentName]; ok {
		return intent(request)
	}

	return alexa.HandleAboutIntent(request)
}

func Handler(request alexa.Request) (alexa.Response, error) {
	return IntentDispatcher(request), nil
}

func main() {
	lambda.Start(Handler)
}
