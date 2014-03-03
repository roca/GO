// tour73
package main

import (
	"fmt"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan map[string]string) {

	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	results_map := <-ch

	if _, present := results_map[url]; depth <= 0 || present {
		ch <- results_map
		return
	}

	var routine = func() {
		body, urls, err := fetcher.Fetch(url)
		if err != nil {

			results_map[fmt.Sprint(err)] = ""
			ch <- results_map
			return
		}

		results_map["found: "+url] = body
		ch <- results_map

		for _, u := range urls {
			Crawl(u, depth-1, fetcher, ch)
		}
	}

	go routine()

	return
}

func main() {

	t1 := time.Now()

	fmt.Printf("Go launched at %s\n", t1)

	x := map[string]string{}

	ch := make(chan map[string]string, 1)
	ch <- x

	Crawl("http://golang.org/", 4, fetcher, ch)

	results := <-ch

	for key, value := range results {
		fmt.Println(key, value)
	}

	t2 := time.Now()

	fmt.Printf("The call took %v to run.\n", t2.Sub(t1))
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
