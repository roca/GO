package cdktable

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	dynamodb "github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
)

type CdktableStackProps struct {
	StackProps awscdk.StackProps
}

func NewCdktableStack(scope constructs.Construct, id string, props *CdktableStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	tableName := "barjokes-cdk"
	
	_,_ = NewDynamoDBTable(&stack, &tableName)

	// example resource
	// queue := awssqs.NewQueue(stack, jsii.String("CdktableQueue"), &awssqs.QueueProps{
	// 	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })

	return stack
}

func NewDynamoDBTable(stack *awscdk.Stack, tableName *string) (*dynamodb.Table, error) {
	params := &dynamodb.TableProps{
		PartitionKey: &dynamodb.Attribute{
			Name: aws.String("NAME"),
			Type: dynamodb.AttributeType_STRING,
		},
		BillingMode: dynamodb.BillingMode_PAY_PER_REQUEST,
		TableName:   tableName,
	}

	table := dynamodb.NewTable(*stack, tableName, params)
	return &table, nil
}
