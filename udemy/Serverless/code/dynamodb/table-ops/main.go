package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var useSamlProfile bool = false

var svc *dynamodb.DynamoDB

// TableName ...
var TableName string = "td_notes"

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	awsConfig := &aws.Config{
		Region:     aws.String("us-east-1"),
		HTTPClient: client,
	}

	if useSamlProfile {
		awsConfig.Credentials = credentials.NewSharedCredentials("", "saml")
		fmt.Println("Using local SAML profile")
	}

	sess, err := session.NewSession(awsConfig)
	if err != nil {
		fmt.Printf("Could not establish session: %v\n", err)
	}

	// Create DynamoDB client
	svc = dynamodb.New(sess)
	//fmt.Println(svc.ClientInfo.ServiceName)
}

func main() {
	// input := &dynamodb.ListTablesInput{}
	// result, err := svc.ListTables(input)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(result)

	// input := &dynamodb.DescribeTableInput{
	// 	TableName: aws.String("td_notes_sdk"),
	// }
	// result, err := svc.DescribeTable(input)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	input := &dynamodb.CreateTableInput{
		TableName: aws.String("td_notes_sdk"),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("user_id"), AttributeType: aws.String("S")},
			{AttributeName: aws.String("timestamp"), AttributeType: aws.String("N")},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("user_id"), KeyType: aws.String("HASH")},
			{AttributeName: aws.String("timestamp"), KeyType: aws.String("RANGE")},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	}
	result, err := svc.CreateTable(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	// input := &dynamodb.UpdateTableInput{
	// 	TableName: aws.String("td_notes_sdk"),
	// 	ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
	// 		ReadCapacityUnits:  aws.Int64(2),
	// 		WriteCapacityUnits: aws.Int64(1),
	// 	},
	// }
	// result, err := svc.UpdateTable(input)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// input := &dynamodb.DeleteTableInput{
	// 	TableName: aws.String("td_notes_sdk"),
	// }
	// result, err := svc.DeleteTable(input)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	str, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	fmt.Println(string(str))

}