package main

import (
	"fmt"
	"log"
	"net/http"
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
}
