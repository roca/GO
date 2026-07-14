package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Input struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

func handler(req *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	input := Input{}
	log.Println(req.Body)
	if err := json.Unmarshal([]byte(req.Body), &input); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response := struct {
		Message   string `json:"message"`
		Timestamp int64  `json:"timestamp"`
	}{
		Message:   fmt.Sprintf("%d + %d =  %d", input.Num1, input.Num2, (input.Num1 + input.Num2)),
		Timestamp: time.Now().Unix(),
	}

	b, _ := json.Marshal(&response)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(b),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
