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

	// input := &dynamodb.GetItemInput{
	// 	TableName: aws.String("td_notes_sdk"),
	// 	Key: map[string]*dynamodb.AttributeValue{
	// 		"user_id":   {S: aws.String("ABC")},
	// 		"timestamp": {N: aws.String("1")},
	// 	},
	// }

	// input := &dynamodb.QueryInput{
	// 	TableName:              aws.String("td_notes_sdk"),
	// 	KeyConditionExpression: aws.String("user_id = :uid"),
	// 	ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
	// 		":uid": {
	// 			S: aws.String("11"),
	// 		},
	// 		":content": {
	// 			S: aws.String("content 11"),
	// 		},
	// 	},
	// 	FilterExpression: aws.String("content = :content"),
	// }

	// input := &dynamodb.ScanInput{
	// 	TableName: aws.String("td_notes"),
	// 	ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
	// 		":cat": {
	// 			S: aws.String("general"),
	// 		},
	// 	},
	// 	FilterExpression: aws.String("cat = :cat"),
	// }

	input := &dynamodb.BatchGetItemInput{
		RequestItems: map[string]*dynamodb.KeysAndAttributes{
			"td_notes_sdk": {
				Keys: []map[string]*dynamodb.AttributeValue{
					{
						"user_id":   {S: aws.String("11")},
						"timestamp": {N: aws.String("1")},
					},
					{
						"user_id":   {S: aws.String("22")},
						"timestamp": {N: aws.String("2")},
					},
				},
			},
			"td_notes": {
				Keys: []map[string]*dynamodb.AttributeValue{
					{
						"user_id":   {S: aws.String("qazwsxf")},
						"timestamp": {N: aws.String("1567571287")},
					},
				},
			},
		},
	}

	result, err := svc.BatchGetItem(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	str, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	fmt.Println(string(str))

}
