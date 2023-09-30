package share

import (
	"bytes"
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var Client *s3.Client

type testClient struct{}

func (tc *testClient) PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {

	return &s3.PutObjectOutput{}, nil

}

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	Client = s3.NewFromConfig(cfg)

}

func UploadToS3(client *s3.Client, fileName *string, bucketName *string) error {

	content, err := os.ReadFile(*fileName)
	if err != nil {
		return err
	}

	_, err = client.PutObject(
		context.TODO(),
		&s3.PutObjectInput{
			Bucket: bucketName,
			Key:    fileName,
			Body:   bytes.NewReader(content),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
