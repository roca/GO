package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/disintegration/imaging"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

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
	dowloadFromS3()

	existingImageFile, err := os.Open("gopher.jpeg")
	if err != nil {
		// Handle error
	}
	defer existingImageFile.Close()

	// Calling the generic image.Decode() will tell give us the data
	// and type of image it is as a string. We expect "png"
	imageData, err := jpeg.Decode(existingImageFile)
	if err != nil {
		// Handle error
	}

	dstImage := imaging.Resize(imageData, 400, 0, imaging.Lanczos)

	// outputFile is a File type which satisfies Writer interface
	outputFile, err := os.Create("test.png")
	if err != nil {
		// Handle error
	}

	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	png.Encode(outputFile, dstImage)

	// Don't forget to close files
	outputFile.Close()

	//lambda.Start(handler)

	uploadToS3()

	os.Remove("gopher.jpeg")
	os.Remove("test.png")
}

func resizeIamge(imageData image.Image) (image.Image, error) {
	return imaging.Resize(imageData, 400, 0, imaging.Lanczos), nil
}

func dowloadFromS3() {
	// NOTE: you need to store your AWS credentials in ~/.aws/credentials

	// 1) Define your bucket and item names
	bucket := "romelbkt"
	item := "gopher.jpeg"

	// 2) Create an AWS session
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	// 3) Create a new AWS S3 downloader
	downloader := s3manager.NewDownloader(sess)

	// 4) Download the item from the bucket. If an error occurs, log it and exit. Otherwise, notify the user that the download succeeded.
	file, err := os.Create(item)
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String("images/" + item),
		})

	if err != nil {
		log.Fatalf("Unable to download item %q, %v", item, err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}

func uploadToS3() {
	// NOTE: you need to store your AWS credentials in ~/.aws/credentials

	// 1) Define your bucket and item names
	bucket := "romelbkt"
	item := "test.png"

	// 2) Create an AWS session
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	// 3) Create a new AWS S3 downloader
	uploader := s3manager.NewUploader(sess)

	// Open the file for use
	file, _ := os.Open(item)
	defer file.Close()

	_, err := uploader.Upload(
		&s3manager.UploadInput{
			Bucket: aws.String(bucket),
			Key:    aws.String("images/dst/" + item),
			Body:   file,
		})

	if err != nil {
		log.Fatalf("Unable to upload item %q, %v", item, err)
	}
}
