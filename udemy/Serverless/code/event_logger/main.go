package main

import (
	"context"
	"log"
	"os"

	"encoding/base64"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

// Event ...
type Event struct {
	Stage          string
	LambdaFunction string
	LambdaVersion  string
}

var encrypted = os.Getenv("APP_SECRET")
var decrypted string

func init() {
	kmsClient := kms.New(session.New())
	decodedBytes, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		panic(err)
	}
	input := &kms.DecryptInput{
		CiphertextBlob: decodedBytes,
	}
	response, err := kmsClient.Decrypt(input)
	if err != nil {
		panic(err)
	}
	// Plaintext is a byte array, so convert to string
	decrypted = string(response.Plaintext[:])
}

func handler(ctx context.Context, event Event) (Event, error) {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Comment
	event.LambdaFunction = lambdacontext.FunctionName
	event.LambdaVersion = lambdacontext.FunctionVersion

	log.Printf("APP_NAME: %s", os.Getenv("APP_NAME"))
	log.Printf("APP_SECRET encrypted: %s", encrypted)
	log.Printf("APP_SECRET decrypted: %s", decrypted)
	return event, nil
}

func main() {
	lambda.Start(handler)
}
