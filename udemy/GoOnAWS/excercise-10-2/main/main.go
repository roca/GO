package main

import (
	share "excercise-10-2"
	"flag"
	"fmt"
	"log"
)

func main() {
	// go run main.go -file testdata/excercise-10-2/L34-exercise-presign.pdf -bucket go-on-aws
	fileName := flag.String("file", "", "The file name")
	bucketName := flag.String("bucket", "", "The bucket name")
	flag.Parse()

	err := share.UploadToS3(share.Client, fileName, bucketName)
	if err != nil {
		log.Fatal(err)
	}

	url, err := share.GetS3Url(share.Client, bucketName, fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
