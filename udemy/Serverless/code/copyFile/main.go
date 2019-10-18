package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
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

	i := 0

	for _, record := range event.Records {
		s3Record := record.S3

		_, err := svc.CopyObject(
			&s3.CopyObjectInput{
				Bucket:     aws.String(os.Getenv("DESTINATION_BUCKET")),
				CopySource: aws.String(s3Record.Bucket.Name + "/" + s3Record.Object.Key),
				Key:        aws.String(s3Record.Object.Key),
			},
		)
		if err != nil {
			return res, err
		}

		event.Records[i].S3.Bucket.Name = os.Getenv("DESTINATION_BUCKET")
		event.Records[i].S3.Object.Key = s3Record.Object.Key
		i++

		b, _ := json.Marshal(event)
		res.StatusCode = http.StatusOK
		res.Body = string(b)

		return res, nil

	}

	return res, errors.New("No file copied")
}

func main() {
	lambda.Start(handler)
}
