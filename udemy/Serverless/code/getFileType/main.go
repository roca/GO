package main

import (
	"context"

	"path/filepath"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sfn"
)

var sess *session.Session
var sf *sfn.SFN

func init() {
	sess = session.Must(session.NewSession())
}

func handler(ctx context.Context, event *events.S3Event) (string, error) {
	for _, record := range event.Records {
		s3 := record.S3
		filename := s3.Object.Key
		suffix := filepath.Ext(filename)
		if suffix != "" {
			return suffix[1:], nil
		}
		return "", nil
	}
	return "", nil
}

func main() {
	lambda.Start(handler)
}
