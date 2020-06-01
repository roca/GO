package main

import (
	"context"
	"errors"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var sess *session.Session
var svc *s3.S3

func init() {
	sess = session.Must(session.NewSession())
	svc = s3.New(sess)
}

type Event struct {
	Records []*events.S3EventRecord `json:"Records"`
	Results struct {
		FileType string `json:"fileType"`
	} `json:"results"`
}

func handler(ctx context.Context, event Event) (bool, error) {

	for _, record := range event.Records {
		s3Record := record.S3

		re := regexp.MustCompile(`\+{1}`)
		key := re.ReplaceAllString(s3Record.Object.Key, " ")
		_, err := svc.DeleteObject(
			&s3.DeleteObjectInput{
				Bucket: aws.String(s3Record.Bucket.Name),
				Key:    aws.String(key),
			},
		)
		if err != nil {
			return false, err
		}
		return true, nil

	}
	return false, errors.New("No file deleted")
}

func main() {
	lambda.Start(handler)
}
