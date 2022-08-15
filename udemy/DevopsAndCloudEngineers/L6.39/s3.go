package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws"
)

func initS3Client(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(regionName), config.WithSharedConfigProfile("go-aws-sdk"))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %s", err)
	}
	return s3.NewFromConfig(cfg), nil
}

func createS3Bucket(ctx context.Context, s3Client *s3.Client) error {
	allBuckets, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return fmt.Errorf("ListBuckets error: %s", err)
	}
	found := false
	for _, bucket := range allBuckets.Buckets {
		if *bucket.Name == bucketName {
			found = true
			break
		}
	}
	if !found {
		_, err = s3Client.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
			CreateBucketConfiguration: &types.CreateBucketConfiguration{
				LocationConstraint: types.BucketLocationConstraint(regionName),
			},
		})
		if err != nil {
			return fmt.Errorf("CreateBucket error: %s", err)
		}
	}

	return nil
}

func upLoadToS3Bucket(ctx context.Context, s3Client *s3.Client) error {
	testFile, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return fmt.Errorf("ReadFile error: %s", err)
	}
	upLoader := manager.NewUploader(s3Client)
	_, err = upLoader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("test.txt"),
		Body:   bytes.NewReader(testFile),
	})
	if err != nil {
		return fmt.Errorf("Upload error: %s", err)
	}
	return nil
}
func downLoadFromS3Bucket(ctx context.Context, s3Client *s3.Client) ([]byte, error) {
	downLoader := manager.NewDownloader(s3Client)
	buffer := manager.NewWriteAtBuffer([]byte{})

	numBytes, err := downLoader.Download(ctx, buffer, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("test.txt"),
	})

	if err != nil {
		return nil, fmt.Errorf("Download error: %s", err)
	}

	if numBytesRecieved := len(buffer.Bytes()); numBytes != int64(numBytesRecieved) {
		return nil, fmt.Errorf("Download error: expected %d bytes, received %d", numBytes, numBytesRecieved)
	}

	return buffer.Bytes(), nil
}
