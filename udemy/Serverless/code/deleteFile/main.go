package main

import (
	"context"
	"errors"
	"os"

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

func handler(ctx context.Context, event *events.S3Event) (bool, error) {

	for _, record := range event.Records {
		s3Record := record.S3

		_, err := svc.DeleteObject(
			&s3.DeleteObjectInput{
				Bucket: aws.String(os.Getenv("DESTINATION_BUCKET")),
				Key:    aws.String(s3Record.Object.Key),
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
