package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var greeting = map[string]string{
	"en": "Hello",
	"fr": "Bonjour",
}

func handler(req *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := req.PathParameters["name"]
	lang := req.QueryStringParameters["lang"]
	infoParams := make(map[string]string)
	for key, value := range req.QueryStringParameters {
		infoParams[key] = value
	}
	delete(infoParams, "lang")

	greetMessage, ok := greeting[lang]
	if !ok {
		greetMessage = greeting["en"]
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       greetMessage + name,
	}, nil
}

func main() {
	lambda.Start(handler)
}
