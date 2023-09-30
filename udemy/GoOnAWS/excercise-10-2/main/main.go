package main

import (
	s3share "excercise-10-2"
	"flag"
	"fmt"
	"log"
)

func main() {
	// go run main.go -file testdata/excercise-10-2/L34-exercise-presign.pdf -bucket go-on-aws
	fileName := flag.String("file", "", "The file name")
	bucketName := flag.String("bucket", "", "The bucket name")
	flag.Parse()

	if *fileName == "" {
		flag.Usage()
		log.Fatal("You must supply a file name")
	}
	if *bucketName == "" {
		flag.Usage()
		log.Fatal("You must supply a bucket name")
	}

	err := s3share.UploadToS3(s3share.Client, fileName, bucketName)
	if err != nil {
		log.Fatal(err)
	}

	url, err := s3share.GetS3Url(s3share.Client, bucketName, fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
