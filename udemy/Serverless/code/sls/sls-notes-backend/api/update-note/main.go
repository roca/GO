package main

/*
	PATCH note/id
*/

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"udemy.com/sls/sls-notes-backend/api/models"
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

	item := models.Item{}

	if err := json.Unmarshal([]byte(event.Body), &item); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	item.UserID = utils.GetUserID(event.Headers)
	item.UserName = utils.GetUserName(event.Headers)
	item.Expires = time.Now().AddDate(0, 0, 90).Unix()

	_, err := svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"user_id":   {S: aws.String(item.UserID)},
			"user_name": {S: aws.String(item.UserName)},
			"expires":   {N: aws.String(fmt.Sprintf("%d", item.Expires))},
		},
		ConditionExpression: aws.String("#t == :t"),
		ExpressionAttributeNames: map[string]*string{
			"#t": aws.String("timestamp"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":t": {
				N: aws.String(fmt.Sprintf("%d", item.TimeStamp)),
			},
		},
	})

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	b, err := json.Marshal(&item)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    utils.GetResponseHeaders(),
		Body:       string(b),
	}

	return response, nil
}

func main() {
	lambda.Start(Handler)
}