package s3share

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetS3Url(client *s3.Client, bucketName *string, key *string) (string, error) {
	lifeTimeSec := int64(3600)
	s3PresignClient := s3.NewPresignClient(client)

	req, err := s3PresignClient.PresignGetObject(
		context.TODO(),
		&s3.GetObjectInput{
			Bucket: bucketName,
			Key:    key,
		},
		func(opts *s3.PresignOptions) {
			opts.Expires = time.Duration(lifeTimeSec * int64(time.Second))
		})

	if err != nil {
		return "", err
	}

	return string(req.URL), nil
}
