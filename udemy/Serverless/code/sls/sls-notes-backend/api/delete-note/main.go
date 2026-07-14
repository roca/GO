package main

/*
	DELETE note/id
*/

import (
	"context"
	"fmt"
	"net/http"

	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"udemy.com/sls/sls-notes-backend/api/utils"
)

var sess *session.Session
var svc *dynamodb.DynamoDB
var tableName string

func init() {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	svc = dynamodb.New(sess)
	tableName = os.Getenv("NOTES_TABLE")
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	pathParams := make(map[string]string)

	for key, value := range event.PathParameters {
		pathParams[key] = value
	}

	//timestamp, err := url.QueryUnescape(pathParams["timestamp"])
	timestamp := pathParams["timestamp"]
	// if err != nil {
	// 	return events.APIGatewayProxyResponse{}, err
	// }

	userID := utils.GetUserID(event.Headers)

	_, err := svc.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"user_id":   {S: aws.String(userID)},
			"timestamp": {N: aws.String(timestamp)},
		},
	})
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    utils.GetResponseHeaders(),
		Body:       fmt.Sprintf("Successfully deleted item for user/timestamp: %s/%s", userID, timestamp),
	}

	return response, nil
}

func main() {
	lambda.Start(Handler)
}
