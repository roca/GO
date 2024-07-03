package main

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func siteTime(url string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error: %s -> %s\n", url, err)
	}

	defer resp.Body.Close()

	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		log.Printf("Error: %s -> %s\n", url, err)
	}

	duration := time.Since(start)
	log.Printf("INFO: %s -> %v\n", url, duration)
}

func main() {

	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.youtube.com",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			siteTime(url)
		}()
	}
	wg.Wait()
}
