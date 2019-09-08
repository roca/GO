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

	// input := &dynamodb.PutItemInput{
	// 	TableName: aws.String("td_notes_sdk"),
	// 	Item: map[string]*dynamodb.AttributeValue{
	// 		"user_id":   {S: aws.String("bb")},
	// 		"timestamp": {N: aws.String("2")},
	// 		"title":     {S: aws.String("changed title")},
	// 		"content":   {S: aws.String("changed content")},
	// 	},
	// }

	// input := &dynamodb.UpdateItemInput{
	// 	TableName: aws.String("td_notes_sdk"),
	// 	ExpressionAttributeNames: map[string]*string{
	// 		"#t": aws.String("title"),
	// 	},
	// 	ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
	// 		":t": {
	// 			S: aws.String("Updated title"),
	// 		},
	// 	},
	// 	Key: map[string]*dynamodb.AttributeValue{
	// 		"user_id":   {S: aws.String("bb")},
	// 		"timestamp": {N: aws.String("1")},
	// 	},
	// 	ReturnValues:     aws.String("UPDATED_NEW"),
	// 	UpdateExpression: aws.String("set #t = :t"),
	// }

	// input := &dynamodb.DeleteItemInput{
	// 	TableName: aws.String("td_notes_sdk"),
	// 	Key: map[string]*dynamodb.AttributeValue{
	// 		"user_id":   {S: aws.String("bb")},
	// 		"timestamp": {N: aws.String("1")},
	// 	},
	// }

	input := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			"td_notes_sdk": {
				{
					DeleteRequest: &dynamodb.DeleteRequest{
						Key: map[string]*dynamodb.AttributeValue{
							"user_id":   {S: aws.String("bb")},
							"timestamp": {N: aws.String("2")},
						},
					},
				},
				{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"user_id":   {S: aws.String("11")},
							"timestamp": {N: aws.String("1")},
							"title":     {S: aws.String("Title 11")},
							"content":   {S: aws.String("content 11")},
						},
					},
				},
				{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"user_id":   {S: aws.String("22")},
							"timestamp": {N: aws.String("2")},
							"title":     {S: aws.String("Title 22")},
							"content":   {S: aws.String("content 22")},
						},
					},
				},
			},
		},
	}

	result, err := svc.BatchWriteItem(input)
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
