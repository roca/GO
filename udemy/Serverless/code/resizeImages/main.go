package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/disintegration/imaging"
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
}

func resizeIamge(imageData image.Image) (image.Image, error) {
	return imaging.Resize(imageData, 400, 0, imaging.Lanczos), nil
}
