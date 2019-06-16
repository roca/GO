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
			"DateID":   {S: aws.String(c.DateID)},
			"Tmstp":    {S: TimetoDB(c.Time)},
			"Username": {S: aws.String(c.Username)},
			"Text":     {S: aws.String(c.Text)},
		},
	})
	return err
}

func GetChat(sess *session.Session) ([]Chat, error) {
	dbc := dynamodb.New(sess)
	// dbres, err := dbc.Query(&dynamodb.QueryInput{
	// 	TableName:              aws.String("ch_chats"),
	// 	KeyConditionExpression: aws.String("DateID = :a"),
	// 	ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
	// 		":a", {S: aws.String(time.Now().Foramt(DATE_FMT))},
	// 	},
	// })

	var queryInput = &dynamodb.QueryInput{
		Limit:     aws.Int64(1),
		TableName: aws.String("ch_chats"),
		KeyConditions: map[string]*dynamodb.Condition{
			"DateID": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(time.Now().Format(DATE_FMT)),
					},
				},
			},
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
