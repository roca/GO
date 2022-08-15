package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	bucketName = "aws-demo-test-bucket-roca"
	regionName = "us-east-2"
)

func main() {
	var (
		s3Client *s3.Client
		err      error
		out      []byte
	)
	ctx := context.Background()
	if s3Client, err = initS3Client(ctx); err != nil {
		fmt.Printf("initS3Client error: %s", err)
		os.Exit(1)
	}
	if err = createS3Bucket(ctx, s3Client); err != nil {
		fmt.Printf("createS3Bucket error: %s", err)
		os.Exit(1)
	}
	if err = upLoadToS3Bucket(ctx, s3Client); err != nil {
		fmt.Printf("upLoadToS3Bucket error: %s", err)
		os.Exit(1)
	}
	fmt.Println("Upload complete.")
	if out, err = downLoadFromS3Bucket(ctx, s3Client); err != nil {
		fmt.Printf("downLoadFromS3Bucket error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("Download complete: %s\n", string(out))
}
