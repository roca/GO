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
	var (
		name = req.PathParameters["name"]
		lang = req.QueryStringParameters["lang"]
	)
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
