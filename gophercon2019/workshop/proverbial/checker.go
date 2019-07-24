package proverbial

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jboursiquot/go-proverbs"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type ddbPutter interface {
	PutItemWithContext(ctx aws.Context, input *dynamodb.PutItemInput, opts ...request.Option) (*dynamodb.PutItemOutput, error)
}

// Checker checks an endpoint to verify it returns a proverb.
type Checker struct {
	httpClient httpClient
	ddbClient  ddbPutter
}

// NewChecker returns a new Checker.
func NewChecker(hc httpClient, ddb ddbPutter) *Checker {
	return &Checker{httpClient: hc, ddbClient: ddb}
}

// Handle handles the request.
func (c *Checker) Handle(ctx context.Context, request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	endpoint, ok := request.QueryStringParameters["endpoint"]
	if !ok {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Endpoint is missing in query string.",
		}, nil
	}

	name, ok := request.QueryStringParameters["name"]
	if !ok {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Name is missing in query string.",
		}, nil
	}

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf("Failed to formulate request to check endpoint: %s", err),
		}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf("Failed to submit request: %s", err),
		}, nil
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf("Failed to read data: %s", err),
		}, nil
	}

	log.Printf("e=%s, n=%s", endpoint, name)

	var p proverbs.Proverb
	if err := json.Unmarshal(data, &p); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf("Failed to unmarshal data: %s", err),
		}, nil
	}

	if p.Saying == "" || p.Link == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotAcceptable,
		}, nil
	}

	chk := check{Endpoint: endpoint, Name: name}
	if err := c.save(ctx, chk); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf("Failed to save check: %s", err),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusAccepted,
	}, nil
}

type check struct {
	ID       string `json:"id"`
	Endpoint string
	Name     string
}

func (c *Checker) save(ctx context.Context, chk check) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	chk.ID = id.String()

	item, err := dynamodbattribute.MarshalMap(chk)
	if err != nil {
		return fmt.Errorf("failed to marshal item for storage: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	}

	if _, err = c.ddbClient.PutItemWithContext(ctx, input); err != nil {
		return fmt.Errorf("failed to save item: %s", err)
	}

	return nil
}
