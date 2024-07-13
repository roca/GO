package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func main() {
	url := "https://d37ci6vzurychx.cloudfront.net/trip-data/yellow_tripdata_2018-05.parquet"
	fileName := "yellow_tripdata_2018-05.parquet"

	// url := "https://www.353solutions.com/c/znga/data/rtb.go"
	// fileName := "rtb.txt"

	//resp, err := http.Head("https://d37ci6vzurychx.cloudfront.net/trip-data/yellow_tripdata_2018-05.parquet")
	resp, err := http.Head(url)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// // Get HEADERS
	for k, v := range resp.Header {
		log.Printf("%v: %v", k, v)
	}
	contentLength := resp.Header.Get("Content-Length")
	eTag := resp.Header.Get("ETag")

	fmt.Printf("Content-Length: %v\n", contentLength)
	fmt.Printf("ETag: %v\n", eTag)

	createEmptyFile(fileName, 0)

	totalSize, err := strconv.Atoi(contentLength)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	n := 10
	chunkSize := totalSize / n
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		offset := i * chunkSize
		size := chunkSize
		if i == n-1 {
			size = totalSize - offset
		}
		wg.Add(1)
		func(url, fileName string, offset, size int) {
			getChunk(url, fileName, offset, size)
			wg.Done()
		}(url, fileName, offset, size)
	}
	wg.Wait()
	fmt.Println("Done")

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
		log.Fatalf("http.NewRequest Error: %v", err)
	}
	// Add the range header to the request
	req.Header.Add("Range", fmt.Sprintf("bytes=%v-%v", offset, offset+size))
	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("http.DefaultClient Error: %v", err)
	}
	// Check if the request was successful
	if resp.StatusCode != http.StatusPartialContent {
		log.Fatalf("Error: %v", resp.Status)
	}
	// fmt.Println("-----------------------------------")
	// for k, v := range resp.Header {
	// 	log.Printf("%v: %v", k, v)
	// }
	// Read the response body

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// fmt.Println("Read", n, "bytes")
	// Write the response body to a file
	err = os.WriteFile(fileName, bytes, 0644)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
