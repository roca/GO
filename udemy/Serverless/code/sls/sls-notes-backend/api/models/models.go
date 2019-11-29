package models

import (
	"strconv"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Item struct {
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
	NoteID    string `json:"note_id"`
	TimeStamp int64  `json:"timestamp"`
	Expires   int64  `json:"expires"`
}

// ExtractItem extracts a DynamoDB record and create a Item object from it
func ExtractItem(record map[string]*dynamodb.AttributeValue) Item {
	var userIDAv, userNameAv, noteIDAv string
	var timeStampAv, expiresAv int64

	if v, ok := record["user_id"]; ok {
		userIDAv = *v.S
	}

	if v, ok := record["user_name"]; ok {
		userNameAv = *v.S
	}

	if v, ok := record["note_id"]; ok {
		noteIDAv = *v.S
	}

	if v, ok := record["timestamp"]; ok {
		timeStampAv, _ = strconv.ParseInt(*v.N, 10, 64)
	}

	if v, ok := record["expires"]; ok {
		expiresAv, _ = strconv.ParseInt(*v.N, 10, 64)
	}

	return Item{
		UserID:    userIDAv,
		UserName:  userNameAv,
		NoteID:    noteIDAv,
		TimeStamp: timeStampAv,
		Expires:   expiresAv,
	}
}
