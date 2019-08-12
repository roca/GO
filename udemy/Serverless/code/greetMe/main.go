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
	if path, ok := req.PathParameters["proxy"]; ok {

	}
	for _, v := range req.QueryStringParameters {
		if greetMessage, ok := greeting[v]; !ok {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       greetMessage + "endpoint does not is exist",
			}, nil
		}

	}

	return events.APIGatewayProxyResponse{}, nil
}

func main() {
	lambda.Start(handler)
}
