package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"udemy.com/sls/sls-notes-backend/api/models"
	uuid "github.com/satori/go.uuid"

)

var sess *session.Session
var svc *kinesis.Kinesis
var streamName *string

func init() {

	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	svc = kinesis.New(sess)
	streamName = aws.String("ServerlessNotesStream")
}

func createRandomNote() string {

	note := models.Note{}
	uuid := uuid.NewV4()

	note.UserID = "KinesisUser"
	note.UserName = "Kinesis User"
	note.NoteID = fmt.Sprintf("%s:%s", note.UserID, uuid)
	note.TimeStamp = time.Now().Unix()
	note.Expires = time.Now().AddDate(0, 0, 90).Unix()
	note.Cat = "general"
	note.Title = "Kinesis TEST"
	note.Content = "Kinesis TEST"


	b, err := json.Marshal(&note)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func main() {

	// put 10 records using PutRecords API
	entries := make([]*kinesis.PutRecordsRequestEntry, 10)
	for i := 0; i < len(entries); i++ {
		entries[i] = &kinesis.PutRecordsRequestEntry{
			Data:         []byte(createRandomNote()),
			PartitionKey: aws.String("key2"),
		}
	}
	fmt.Printf("%v\n", entries)
	putsOutput, err := svc.PutRecords(&kinesis.PutRecordsInput{
		Records:    entries,
		StreamName: streamName,
	})
	if err != nil {
		panic(err)
	}
	// putsOutput has Records, and its shard id and sequece enumber.
	fmt.Printf("%v\n", putsOutput)

}
