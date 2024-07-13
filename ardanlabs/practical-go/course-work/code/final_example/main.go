package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	resp, err := http.Head("https://d37ci6vzurychx.cloudfront.net/trip-data/yellow_tripdata_2018-05.parquet")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// // Get HEADERS
	// for k, v := range resp.Header {
	// 	log.Printf("%v: %v", k, v)
	// }
	contentLength := resp.Header.Get("Content-Length")
	eTag := resp.Header.Get("ETag")

	fmt.Printf("Content-Length: %v\n", contentLength)
	fmt.Printf("ETag: %v\n", eTag)

	fileName := "yellow_tripdata_2018-05.parquet"
	createEmptyFile(fileName, 0)

	totalSize,err := strconv.Atoi(contentLength)
	chunkSize := totalSize / 10

	for i := 0; i < 10; i++ {
		getChunk("https://d37ci6vzurychx.cloudfront.net/trip-data/yellow_tripdata_2018-05.parquet", fileName, i*chunkSize, chunkSize)
	}

}

func createEmptyFile(path string, size int64) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Seek(size-1, io.SeekStart)
	file.Write([]byte{0})
	return nil
}

func getChunk(url, fileName string, offset, size int) {
	// Create a new HEAD request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// Add the range header to the request
	req.Header.Add("Range", fmt.Sprintf("bytes=%v-%v", offset, offset+size))
	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// Check if the request was successful
	if resp.StatusCode != http.StatusPartialContent {
		log.Fatalf("Error: %v", resp.Status)
	}
	// Read the response body
	buf := make([]byte, size)
	_, err = resp.Body.Read(buf)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// Write the response body to a file
	err = os.WriteFile(fileName, buf, 0644)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
