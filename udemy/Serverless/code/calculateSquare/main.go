package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	svclambda "github.com/aws/aws-sdk-go/service/lambda"
)

var sess *session.Session

func init() {
	sess = session.Must(session.NewSession())

}

type Event struct {
	Number int
}

type svcEvent struct {
	Operation string `json:"Operation"`
	Input     struct {
		Operand1 int `json:"Operand1"`
		Operand2 int `json:"Operand2"`
	} `json:"Input"`
}

func handler(ctx context.Context, event Event) (events.APIGatewayProxyResponse, error) {
	svc := svclambda.New(sess)

	payload := svcEvent{Operation: "multiply"}
	payload.Input.Operand1 = event.Number
	payload.Input.Operand2 = event.Number

	b, _ := json.Marshal(&payload)

	input := &svclambda.InvokeInput{
		//ClientContext: ctx,
		FunctionName:   aws.String("calculator"),
		InvocationType: aws.String("RequestResponse"),
		Payload:        b,
	}

	res := events.APIGatewayProxyResponse{}

	result, err := svc.Invoke(input)
	if err != nil {
		return res, err
	}

	if err := json.Unmarshal(result.Payload, &res); err != nil {
		return res, err
	}

	log.Println(event)
	log.Println(result)
	return res, nil
}

func main() {
	lambda.Start(handler)
}
