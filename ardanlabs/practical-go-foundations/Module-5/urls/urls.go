package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://go.dev",
		"https://ardanlabs.com",
		"https://ibm.com/no/such/page",
	}

	start := time.Now()
	for _, url := range urls {
		stat, err := urlCheck(url)
		fmt.Printf("%q: %d (%v)\n", url, stat, err)
	}
	duration := time.Since(start)
	fmt.Printf("%d urls in %v\n", len(urls), duration)
	fmt.Println("-------------fanOut----------------")
	fanOutResult(urls)
	fmt.Println("-------------fanOutWait----------------")
	fanOutWait(urls)
	fmt.Println("-------------fanOutPool----------------")
	fanOutPool(urls)
}

func urlLog(url string) {
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("urlLog", "url", url, "error", err)
		return
	}
	slog.Info("urlLog", "url", url, "status", resp.StatusCode)
}

func fanOutPool(urls []string) {
	start := time.Now()
	var wg sync.WaitGroup

	ch := make(chan string)

	go func() { // Producer
		for _, url := range urls {
			ch <- url
		}
		close(ch)
	}()

	const size = 2
	wg.Add(size)
	for range size { // Consumers
		go func() {
			defer wg.Done()
			for url := range ch {
				urlLog(url)
			}
		}()
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("%d urls in %v\n", len(urls), duration)
}

func fanOutWait(urls []string) {
	start := time.Now()

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			urlLog(url)
		}()
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("%d urls in %v\n", len(urls), duration)
}

func fanOutResult(urls []string) {
	start := time.Now()
	type result struct {
		url    string
		status int
		err    error
	}
	ch := make(chan result)
	for _, url := range urls {
		go func() {
			r := result{url: url}
			defer func() {
				ch <- r
			}()

			r.status, r.err = urlCheck(url)
		}()
	}

	for range urls {
		r := <-ch
		fmt.Printf("%q: %d (%v)\n", r.url, r.status, r.err)
	}
	duration := time.Since(start)
	fmt.Printf("%d urls in %v\n", len(urls), duration)
}

func urlCheck(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}
