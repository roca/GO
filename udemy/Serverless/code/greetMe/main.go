package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var greeting = map[string]string{
	"en": "Hello",
	"fr": "Bonjour",
}

func handler(req *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := req.PathParameters["name"]
	infoParams := make(map[string]string)

	for key, value := range req.QueryStringParameters {
		infoParams[key] = value
	}
	lang := infoParams["lang"]
	delete(infoParams, "lang")

	greetMessage, ok := greeting[lang]
	if !ok {
		greetMessage = greeting["en"]
	}

	message := fmt.Sprintf("%s %s", greetMessage, name)
	response := struct {
		Message   string            `json:"message"`
		Info      map[string]string `json:"info"`
		Timestamp int64             `json:"timestamp"`
	}{
		Message:   message,
		Info:      infoParams,
		Timestamp: time.Now().Unix(),
	}

	b, _ := json.Marshal(&response)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(b),
	}, nil
}

func main() {
	lambda.Start(handler)
}
