package main

/*
	Route: POST /note
*/

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"encoding/json"
	"os"

	uuid "github.com/satori/go.uuid"

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
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	svc = dynamodb.New(sess)
	tableName = os.Getenv("NOTES_TABLE")
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	note := models.Note{}
	uuid := uuid.NewV4()

	if err := json.Unmarshal([]byte(event.Body), &note); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	note.UserID = utils.GetUserID(event.Headers)
	note.UserName = utils.GetUserName(event.Headers)
	note.NoteID = fmt.Sprintf("%s:%s", note.UserID, uuid)
	note.TimeStamp = time.Now().Unix()
	note.Expires = time.Now().AddDate(0, 0, 90).Unix()

	_, err := svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"user_id":   {S: aws.String(note.UserID)},
			"user_name": {S: aws.String(note.UserName)},
			"note_id":   {S: aws.String(note.NoteID)},
			"timestamp": {N: aws.String(fmt.Sprintf("%d", note.TimeStamp))},
			"expires":   {N: aws.String(fmt.Sprintf("%d", note.Expires))},
			"cat":       {S: aws.String(note.Cat)},
			"title":     {S: aws.String(note.Title)},
			"content":   {S: aws.String(note.Content)},
		},
	})

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
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

func main() {
	lambda.Start(Handler)
}
