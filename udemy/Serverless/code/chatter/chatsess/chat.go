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
	var dateav, unameav, txav string
	var timeav time.Time

	if v, ok := item["DateID"]; ok {
		dateav = *v.S
	}

	if v, ok := item["Tmstp"]; ok {
		timeav = DBtoTime(v.N)
	}

	if v, ok := item["Username"]; ok {
		unameav = *v.S
	}

	if v, ok := item["Text"]; ok {
		txav = *v.S
	}

	return Chat{
		DateID:   dateav,
		Time:     timeav,
		Username: unameav,
		Text:     txav,
	}
}

// Put ..
func (c Chat) Put(sess *session.Session) error {
	dbc := dynamodb.New(sess)
	_, err := dbc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("ch_chats"),
		Item: map[string]*dynamodb.AttributeValue{
			"DateID":   {S: aws.String(c.DateID)},
			"Tmstp":    {N: TimetoDB(c.Time)},
			"Username": {S: aws.String(c.Username)},
			"Text":     {S: aws.String(c.Text)},
		},
	})
	return err
}

// GetChat ...
func GetChat(sess *session.Session) ([]Chat, error) {
	dbc := dynamodb.New(sess)

	var queryInput = &dynamodb.QueryInput{
		TableName:              aws.String("ch_chats"),
		KeyConditionExpression: aws.String("DateID = :a"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":a": {S: aws.String(time.Now().Format(DATE_FMT))},
		},
		// KeyConditions: map[string]*dynamodb.Condition{
		// 	"DateID": {
		// 		ComparisonOperator: aws.String("EQ"),
		// 		AttributeValueList: []*dynamodb.AttributeValue{
		// 			{
		// 				S: aws.String(time.Now().Format(DATE_FMT)),
		// 			},
		// 		},
		// 	},
		// },
	}

	dbres, err := dbc.Query(queryInput)

	if err != nil {
		return []Chat{}, err
	}

	res := []Chat{}
	for _, v := range dbres.Items {
		res = append(res, ChatFromItem(v))
	}

	return res, nil

}

// GetChatAfter ...
func GetChatAfter(DateID string, t time.Time, sess *session.Session) ([]Chat, error) {
	dbc := dynamodb.New(sess)
	var queryInput = &dynamodb.QueryInput{
		TableName:              aws.String("ch_chats"),
		KeyConditionExpression: aws.String("DateID = :a"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":a": {S: aws.String(time.Now().Format(DATE_FMT))},
		},
		// KeyConditions: map[string]*dynamodb.Condition{
		// 	"DateID": {
		// 		ComparisonOperator: aws.String("EQ"),
		// 		AttributeValueList: []*dynamodb.AttributeValue{
		// 			{
		// 				S: aws.String(time.Now().Format(DATE_FMT)),
		// 			},
		// 		},
		// 	},
		// },
		ExclusiveStartKey: map[string]*dynamodb.AttributeValue{
			"DateID": {S: aws.String(DateID)},
			"Tmstp":  {N: TimetoDB(t)},
		},
	}

	dbres, err := dbc.Query(queryInput)

	if err != nil {
		return []Chat{}, err
	}

	res := []Chat{}
	for _, v := range dbres.Items {
		res = append(res, ChatFromItem(v))
	}

	return res, nil

}
