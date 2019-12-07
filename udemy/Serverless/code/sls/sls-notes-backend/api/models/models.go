package models

import (
	"strconv"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Note struct {
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
	NoteID    string `json:"note_id"`
	TimeStamp int64  `json:"timestamp"`
	Expires   int64  `json:"expires"`
	Cat       string `json:"cat"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

type LastEvaluatedKey struct {
	UserID    string `json:"user_id"`
	TimeStamp int64  `json:"timestamp"`
}

type Notes struct {
	Notes            []Note `json:"notes"`
	Count            int64  `json:"Count"`
	ScannedCount     int64  `json:"ScannerCount"`
	LastEvaluatedKey `json:"LastEvaluatedKey"`
}

// ExtractNotes extracts a Notes object from DynamoDB QueryOutput
func ExtractNotes(data *dynamodb.QueryOutput) Notes {
	timestamp, _ := strconv.ParseInt(*(data.LastEvaluatedKey["timestamp"]).S, 10, 64)
	notes := Notes{
		Count:        *(data.Count),
		ScannedCount: *(data.ScannedCount),
		LastEvaluatedKey: LastEvaluatedKey{
			UserID:    *(data.LastEvaluatedKey["user_id"]).S,
			TimeStamp: timestamp,
		},
	}

	items := []Note{}
	for _, v := range data.Items {
		items = append(items, ExtractNote(v))
	}
	notes.Notes = items

	return notes
}

// ExtractNote extracts a DynamoDB record and create a Note object from it
func ExtractNote(record map[string]*dynamodb.AttributeValue) Note {
	var userIDAv, userNameAv, noteIDAv, catAv, titleAv, contentAv string
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

	if v, ok := record["cat"]; ok {
		catAv = *v.S
	}

	if v, ok := record["title"]; ok {
		titleAv = *v.S
	}

	if v, ok := record["content"]; ok {
		contentAv = *v.S
	}

	return Note{
		UserID:    userIDAv,
		UserName:  userNameAv,
		NoteID:    noteIDAv,
		TimeStamp: timeStampAv,
		Expires:   expiresAv,
		Cat:       catAv,
		Title:     titleAv,
		Content:   contentAv,
	}
}
