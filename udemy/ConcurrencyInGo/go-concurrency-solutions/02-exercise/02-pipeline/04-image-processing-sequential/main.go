package main

import (
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/disintegration/imaging"
)

// Image processing - sequential
// Input - directory with images.
// output - thumbnail images

// Pipeline
// walkfile ----------> processImage -----------> saveImage
//            (paths)                  (results)

type result struct {
	srcImagePath   string
	thumbnailImage *image.NRGBA
	err            error
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("need to send directory path of images")
	}
	start := time.Now()
	err := setupPipeline(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Time taken: %s\n", time.Since(start))
}

func setupPipeline(root string) error {
	done := make(chan struct{})
	defer close(done)

	// first stage of pipline
	paths, errc := walkFiles(done, root)

	// second stage of pipline
	// process the image
	results := processImage(done, paths)

	// third stage of pipline
	for r := range results {
		if r.err != nil {
			return r.err
		}
		saveThumbnail(r.srcImagePath, r.thumbnailImage)
	}

	// check for error on the channel, from walkfiles stage.
	if err := <-errc; err != nil {
		return err
	}
	return nil
}

func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)
	go func() {
		defer close(paths)
		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

			// filter out error
			if err != nil {
				return err
			}

			// check if it is file
			if !info.Mode().IsRegular() {
				return nil
			}

			// check if it is image/jpeg
			contentType, _ := getFileContentType(path)
			if contentType != "image/jpeg" {
				return nil
			}
			select {
			case paths <- path:
			case <-done:
				return fmt.Errorf("walk was cancelled")
			}
			return nil
		})

	}()
	return paths, errc
}

func processImage(done <-chan struct{}, paths <-chan string) <-chan *result {
	results := make(chan *result)
	var wg sync.WaitGroup

	thumbnailer := func() {
		for path := range paths {
			// load the image from file
			srcImage, err := imaging.Open(path)
			if err != nil {
				select {
				case results <- &result{path, nil, err}:
				case <-done:
					return
				}
			}

			// scale the image to 100px * 100px
			thumbnailImage := imaging.Thumbnail(srcImage, 100, 100, imaging.Lanczos)

			select {
			case results <- &result{path, thumbnailImage, nil}:
			case <-done:
				return
			}

		}
	}
	const numThumbnailer = 5
	for i := 0; i < numThumbnailer; i++ {
		wg.Add(1)
		go func() {
			thumbnailer()
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()
	return results
}

// saveThumbnail - save the thumnail image to folder
func saveThumbnail(srcImagePath string, thumbnailImage *image.NRGBA) error {
	filename := filepath.Base(srcImagePath)
	dstImagePath := "thumbnail/" + filename

	// save the image in the thumbnail folder.
	err := imaging.Save(thumbnailImage, dstImagePath)
	if err != nil {
		return err
	}
	fmt.Printf("%s -> %s\n", srcImagePath, dstImagePath)
	return nil
}

// getFileContentType - return content type and error status
func getFileContentType(file string) (string, error) {

	out, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err = out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
