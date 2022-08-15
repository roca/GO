package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type IS3Client interface {
	ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
	CreateBucket(ctx context.Context, params *s3.CreateBucketInput, optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error)
}

type IS3Uploader interface {
	Upload(ctx context.Context, input *s3.PutObjectInput, opts ...func(*manager.Uploader)) (*manager.UploadOutput, error)
}

type IS3Downloader interface {
	Download(ctx context.Context, w io.WriterAt, input *s3.GetObjectInput, options ...func(*manager.Downloader)) (n int64, err error)
}

func initS3Client(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(regionName), config.WithSharedConfigProfile("go-aws-sdk"))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %s", err)
	}
	return s3.NewFromConfig(cfg), nil
}

func createS3Bucket(ctx context.Context, s3Client IS3Client) error {
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

func upLoadToS3Bucket(ctx context.Context, upLoader IS3Uploader, filename string) error {
	testFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("ReadFile error: %s", err)
	}
	_, err = upLoader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(testFile),
	})
	if err != nil {
		return fmt.Errorf("Upload error: %s", err)
	}
	return nil
}
func downLoadFromS3Bucket(ctx context.Context, downLoader IS3Downloader) ([]byte, error) {
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
