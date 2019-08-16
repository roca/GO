package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"

	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/disintegration/imaging"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func handler(s3Event *events.S3Event) (string, error) {
	for _, record := range s3Event.Records {
		s3 := record.S3
		message := fmt.Sprintf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
		Info.Println("Special Information" + message)
		if s3.Bucket.Name == "romelbkt" && s3.Object.Key == "images/gopher.jpeg" {
			resizeImage()
		}
	}

	return "", nil
}

func resizeImage() {
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

	// Get file from s3
	// Resize the file
	// Read the new resized file
	// Upload the new files to s3
}

func main() {
	initLoggers(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	lambda.Start(handler)

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

func initLoggers(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
