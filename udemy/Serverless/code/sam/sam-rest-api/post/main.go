package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var sess *session.Session
var svc *dynamodb.DynamoDB

func init() {
	sess = session.Must(session.NewSession())
	svc = dynamodb.New(sess)
}

type Item struct {
	UserID    string `json:"userid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Website   string `json:"website"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	userid := request.PathParameters["userid"]

	item := Item{}

	_ = json.Unmarshal([]byte(request.Body), &item)
	item.UserID = userid

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello", ""),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
