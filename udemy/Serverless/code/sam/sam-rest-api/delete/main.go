package main

import (
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

	err := Delete(userid)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       "User deleted successfully",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

// Delete deletes a user from the database
func Delete(userID string) error {

	_, err := svc.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"userid": {S: aws.String(userID)},
		},
	})
	if err != nil {
		return err
	}

	return nil

}