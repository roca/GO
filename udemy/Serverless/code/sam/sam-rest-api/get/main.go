package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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

	item, err := Get(userid)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
                Body:       err.Error(),
		}, nil
	}

	b, _ := json.Marshal(&item)

	return events.APIGatewayProxyResponse{
		Body:       string(b),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

// Get Gets a user from the database
func Get(userID string) (Item, error) {

	item := Item{}

	record, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"userid": {S: aws.String(userID)},
		},
	})
	if err != nil {
		return item, err
	}
	if record.Item == nil {
		return item, fmt.Errorf("No User exists by that name :%s", userID)
	}

	firstname := record.Item["firstname"]
	lastname := record.Item["lastname"]
	email := record.Item["email"]
	website := record.Item["website"]

	return Item{
		UserID:    userID,
		FirstName: *(firstname.S),
		LastName:  *(lastname.S),
		Email:     *(email.S),
		Website:   *(website.S),
	}, nil

}
