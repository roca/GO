package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var sess *session.Session
var svc *dynamodb.DynamoDB
var tableName = os.Getenv("TABLE_NAME")

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

	if err := json.Unmarshal([]byte(request.Body), &item); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	item.UserID = userid

	if err := Put(item); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       "Data inserted/updated successfully",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

// Put into Database
func Put(item Item) error {

	_, err := svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"timestamp": {S: aws.String(fmt.Sprintf("%d", time.Now().UnixNano()))},
			"userid":    {S: aws.String(item.UserID)},
			"firstname": {S: aws.String(item.FirstName)},
			"lastname":  {S: aws.String(item.LastName)},
			"email":     {S: aws.String(item.Email)},
			"website":   {S: aws.String(item.Website)},
		},
	})

	return err
}
