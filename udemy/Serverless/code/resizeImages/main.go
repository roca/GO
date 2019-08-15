package main

import "github.com/aws/aws-lambda-go/lambda"

type Event struct {
}

func handler(e Event) (string, error) {

	// Get file from s3
	// Resize the file
	// Read the new resized file
	// Upload the new files to s3
	return "", nil
}

func main() {
	lambda.Start(handler)
}
