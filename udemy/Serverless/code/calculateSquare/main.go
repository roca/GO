package main

import (
	"context"
	"encoding/json"
	"log"

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

type Result struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       struct {
		Answer int `json:"answer"`
	} `json:"body"`
}

type Response struct {
	Answer int `json:"answer"`
}

func handler(ctx context.Context, event Event) (Response, error) {
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

	res := Response{}
	result := Result{}

	rslt, err := svc.Invoke(input)
	if err != nil {
		return res, err
	}

	if err := json.Unmarshal(rslt.Payload, &result); err != nil {
		return res, err
	}

	res.Answer = result.Body.Answer

	log.Println(event)
	log.Println(result)
	return res, nil
}

func main() {
	lambda.Start(handler)
}
