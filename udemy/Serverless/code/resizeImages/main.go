package main

import (
	"image/jpeg"
	"image/png"
	"log"
	"regexp"
	"strings"
	"time"

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

var sess *session.Session

func init() {
	sess = session.Must(session.NewSession())
}

func handler(s3Event *events.S3Event) (string, error) {
	for _, record := range s3Event.Records {
		s3 := record.S3
		message := fmt.Sprintf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
		Info.Println("Special Information" + message)
		matchedKey, _ := regexp.Match(`images/.*\.jpeg`, []byte(s3.Object.Key))
		if s3.Bucket.Name == "romelbkt" && matchedKey {
			err := resizeImage(s3.Bucket.Name, s3.Object.Key)
			if err != nil {
				return "", err
			}
		}
	}
	Info.Println("Image successfully resized")
	return "Image successfully resized", nil
}

func main() {
	initLoggers(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	lambda.Start(handler)

}

func resizeImage(bucketName string, key string) error {
	// Get file from s3
	existingImageFile, err := dowloadFromS3(bucketName, key)
	if err != nil {
		return err
	}

	// Calling the generic image.Decode() will tell give us the data
	// and type of image it is as a string. We expect "png"
	imageData, err := jpeg.Decode(existingImageFile)
	if err != nil {
		return err
	}

	dstImage := imaging.Resize(imageData, 400, 0, imaging.Lanczos)

	newImageFileName := fmt.Sprintf("%d.png", time.Now().UnixNano())

	// outputFile is a File type which satisfies Writer interface
	// TODO: Make file name unique with uuid
	outputFile, err := os.Create("/tmp/" + newImageFileName)
	if err != nil {
		Error.Println("Could create tmp file:", err)
		return err
	}

	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	png.Encode(outputFile, dstImage)

	// Don't forget to close files
	outputFile.Close()

	// Upload the new files to s3
	err = uploadToS3(newImageFileName, bucketName, "images/dst/"+newImageFileName)
	if err != nil {
		Error.Println("Could upload image to s3:", err)
		return err
	}

	os.Remove("/tmp/" + strings.Split(key, "/")[1])
	os.Remove("/tmp/" + newImageFileName)
	Info.Println("Successfully resized image:", key)
	return nil
}

func dowloadFromS3(bucketName string, key string) (*os.File, error) {
	// Create a new AWS S3 downloader
	downloader := s3manager.NewDownloader(sess)

	// 4) Download the item from the bucket. If an error occurs, log it and exit.
	// Otherwise, notify the user that the download succeeded.
	// TODO: Make file name unique with uuid
	file, err := os.Create("/tmp/" + strings.Split(key, "/")[1])
	if err != nil {
		Error.Println("Could not create tmp file:", strings.Split(key, "/")[1])
		return &os.File{}, err
	}

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(key),
		})
	if err != nil {
		Error.Println("Could not download file:", key)
		return &os.File{}, err
	}

	Info.Println("Successfully Downloaded", file.Name(), numBytes, "bytes")
	return file, nil
}

func uploadToS3(fileName string, bucketName string, key string) error {

	file, err := os.Open("/tmp/" + fileName)
	if err != nil {
		Error.Println("Could open tmp file:", err)
		return err
	}
	defer file.Close()

	// Create a new AWS S3 uploader
	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(
		&s3manager.UploadInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(key),
			Body:   file,
		})

	if err != nil {
		Error.Println("Could not upload file:", key)
		return err
	}
	Info.Println("Successfully uploaded to S3:", key)
	return nil
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
