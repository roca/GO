package proverbial

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type ddbScanner interface {
	ScanWithContext(ctx aws.Context, input *dynamodb.ScanInput, opts ...request.Option) (*dynamodb.ScanOutput, error)
}

// Picker picks an item from the checks database.
type Picker struct {
	ddbClient ddbScanner
}

// NewPicker returns a new Picker.
func NewPicker(ddb ddbScanner) *Picker {
	return &Picker{ddbClient: ddb}
}

// Handle handles the request.
func (p *Picker) Handle(ctx context.Context) (string, error) {
	checks, err := p.getAllItems(ctx)
	if err != nil {
		return "", err
	}

	// seed our random picking and pick
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	// pick
	return checks[r.Intn(len(checks))].Name, nil
}

func (p *Picker) getAllItems(ctx context.Context) ([]check, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	}
	result, err := p.ddbClient.ScanWithContext(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to scan table: %s", err)
	}

	var checks []check
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &checks)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal results: %s", err)
	}

	return checks, nil
}
