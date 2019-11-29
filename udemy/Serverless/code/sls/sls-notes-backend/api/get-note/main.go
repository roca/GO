package main

/*
	GET /note
	GET /note/{note_id}
*/

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

	queryParams := make(map[string]string)
	pathParams := make(map[string]string)

	queryParams["limit"] = "5"
	queryParams["start"] = "0"
	for key, value := range event.QueryStringParameters {
		queryParams[key] = value
	}
	for key, value := range event.PathParameters {
		pathParams[key] = value
	}

	limit, _ := strconv.ParseInt(queryParams["limit"], 10, 64)
	startTimeStamp, _ := strconv.ParseInt(queryParams["start"], 10, 64)
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

	if v, ok := pathParams["note_id"]; ok {
		decodedValue, err := url.QueryUnescape(v)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
		queryInput.IndexName = aws.String("note_id-index")
		queryInput.KeyConditionExpression = aws.String("note_id= :id")
		queryInput.ExpressionAttributeValues = map[string]*dynamodb.AttributeValue{
			":id": {S: aws.String(decodedValue)},
		}
		queryInput.Limit = aws.Int64(1)

		data, err := svc.Query(&queryInput)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}

		if len(data.Items) == 0 {
			response := events.APIGatewayProxyResponse{
				StatusCode: http.StatusFound,
				Headers:    utils.GetResponseHeaders(),
			}
			return response, nil
		}

		note := models.Note{}
		for _, v := range data.Items {
			note = models.ExtractNote(v)
			break
		}

		b, err := json.Marshal(&note)
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

	data, err := svc.Query(&queryInput)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	notes := []models.Note{}
	for _, v := range data.Items {
		notes = append(notes, models.ExtractNote(v))
	}

	b, err := json.Marshal(&notes)
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
