package main

/*
	GET note
	GET note/id
*/

import (
	"context"
	"net/http"

	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"udemy.com/sls/sls-notes-backend/api/utils"
)

var sess *session.Session
var svc *dynamodb.DynamoDB
var tableName string

func init() {
	sess = session.Must(session.NewSession())
	svc = dynamodb.New(sess)
	tableName = os.Getenv("NOTES_TABLE")
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	response := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    utils.GetResponseHeaders(),
		Body:       "Success",
	}

	return response, nil
}

func main() {
	lambda.Start(Handler)
}
