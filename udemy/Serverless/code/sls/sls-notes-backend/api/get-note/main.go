package main

/*
	GET /note
*/

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

	infoParams := make(map[string]string)

	infoParams["limit"] = "5"
	infoParams["start"] = "0"
	for key, value := range event.QueryStringParameters {
		infoParams[key] = value
	}

	limit, _ := strconv.ParseInt(infoParams["limit"], 10, 64)
	startTimeStamp, _ := strconv.ParseInt(infoParams["start"], 10, 64)
	userID := utils.GetUserID(event.Headers)

	queryInput := dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("user_id= :id"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {S: aws.String(userID)},
		},
		Limit:            aws.Int64(limit),
		ScanIndexForward: aws.Bool(false),
	}

	if startTimeStamp > 0 {
		queryInput.ExclusiveStartKey = map[string]*dynamodb.AttributeValue{
			"user_id":   {S: aws.String(userID)},
			"timestamp": {N: aws.String(fmt.Sprintf("%d", startTimeStamp))},
		}
	}

	records, err := svc.Query(&queryInput)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	items := []models.Item{}
	for _, v := range records.Items {
		items = append(items, models.ExtractItem(v))
	}

	b, err := json.Marshal(&items)
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
