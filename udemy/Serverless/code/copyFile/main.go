package main

import (
	"context"
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

func handler(ctx context.Context, event *events.S3Event) (events.APIGatewayProxyResponse, error) {

	res := events.APIGatewayProxyResponse{}

	for _, record := range event.Records {
		s3 := record.S3

		_, err := svc.CopyObject(
			&s3.CopyObjectInput{
				Bucket:     aws.String(os.Getenv("DESTINATION_BUCKET")),
				CopySource: aws.String(s3.Bucket.Name),
				Key:        aws.String(item),
			},
		)
		if err != nil {

		}
	}

	return res, nil
}

func main() {
	lambda.Start(handler)
}
