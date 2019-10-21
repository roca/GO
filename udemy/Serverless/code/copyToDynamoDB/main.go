package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var sess *session.Session
var svc *dynamodb.DynamoDB

// Event ...
type Event []byte

func handler(ctx context.Context, event Event) (string, error) {

	images := make(map[string]interface{})

	// unmarschal JSON
	e := json.Unmarshal(event, &images)

	// panic on error
	if e != nil {
		panic(e)
	}

	// a string slice to hold the keys
	k := make([]string, len(images))

	// iteration counter
	i := 0

	// copy c's keys into k
	for s, _ := range images {
		k[i] = s
		i++
	}

	log.Println(k)

	return "Hello World", nil
}

func main() {
	lambda.Start(handler)
}
