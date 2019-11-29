package main

/*
	GET /note
*/

import (
	"context"
	"net/http"
	"strconv"

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
	sess = session.Must(session.NewSession())
	svc = dynamodb.New(sess)
	tableName = os.Getenv("NOTES_TABLE")
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	infoParams := make(map[string]string)

	infoParams["limit"] = "5"
	for key, value := range event.QueryStringParameters {
		infoParams[key] = value
	}

	limit, _ := strconv.ParseInt(infoParams["limit"], 10, 64)
	userID := utils.GetUserID(event.Headers)

	_, err := svc.Query(&dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("user_id= :id"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {S: aws.String(userID)},
		},
		Limit:            aws.Int64(limit),
		ScanIndexForward: aws.Bool(false),
	})
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

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
