package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sfn"
)

var sess *session.Session
var sf *sfn.SFN

func init() {
	sess = session.Must(session.NewSession())
	sf = sfn.New(sess)
}

func handler(ctx context.Context, event *events.S3Event) (events.APIGatewayProxyResponse, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	b, _ := json.Marshal(event)
	res := events.APIGatewayProxyResponse{}

	log.Printf("Processing %s", lambdacontext.FunctionName)

	in := sfn.StartExecutionInput{
		StateMachineArn: aws.String(os.Getenv("STATE_MACHINE_ARN")),
		Input:           aws.String(string(b)),
	}

	out, err := sf.StartExecutionWithContext(ctx, &in)
	if err != nil {
		return res, err
	}

	log.Printf("Started State Machine execution | ARN: %s", *out.ExecutionArn)

	res.StatusCode = http.StatusOK
	res.Body = string(b)

	return res, nil
}

func main() {
	lambda.Start(handler)
}
