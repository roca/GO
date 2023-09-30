package s3share

import (
	"bytes"
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadToS3(client S3Client, fileName *string, bucketName *string) error {

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
