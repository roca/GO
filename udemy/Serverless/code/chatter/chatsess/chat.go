package chatsess

import (
	"html"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Chat ...
type Chat struct {
	DateID   string
	Time     time.Time
	Username string
	Text     string
}

// NewChat ..
func NewChat(Username, Text string) Chat {
	return Chat{
		DateID:   time.Now().Format(DATE_FMT),
		Time:     time.Now(),
		Username: Username,
		Text:     html.EscapeString(Text),
	}
}

// ChatFromItem ...
func ChatFromItem(item map[string]*dynamodb.AttributeValue) Chat {
	chat := Chat{}
	if dateav, ok := item["DateID"]; ok {
		chat.DateID = *dateav.S
	}

	if timeav, ok := item["Tmstp"]; ok {
		chat.Time = DBtoTime(timeav.N)
	}

	if unameav, ok := item["Username"]; ok {
		chat.Username = *unameav.S
	}

	if txav, ok := item["Username"]; ok {
		chat.Text = *txav.S
	}

	return chat
}

// Put ..
func (c Chat) Put(sess *session.Session) error {
	dbc := dynamodb.New(sess)
	_, err := dbc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("ch_chats"),
		Item: map[string]*dynamodb.AttributeValue{
			"DateID":   aws.String(c.DateID),
			"Tmstp":    TimetoDB(c.Time),
			"Username": aws.String(c.Username),
			"Text":     aws.String(c.Text),
		},
	})
	return err
}
