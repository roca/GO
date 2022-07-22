package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// {"page":"words","input":"word1","words":["word1"]}
type Page struct {
	Name string `json:"page"`
}

type Words struct {
	Input string   `json:"input"`
	Words []string `json:"words"`
}

type Occurrence struct {
	Words map[string]int `json:"words"`
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage: ./http-get <url>n")
		os.Exit(1)
	}

	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("URL is in invalid format: %s\n", err)
		os.Exit(1)
	}

	response, err := http.Get(args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		fmt.Printf("Invalid output (HTTP Code %d): %sn", response.StatusCode, body)
		os.Exit(1)
	}

	var page Page
	if err := json.Unmarshal(body, &page); err != nil {
		log.Fatal(err)
	}

	switch page.Name {
	case "words":
		var words Words
		if err := json.Unmarshal(body, &words); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("JSON Parsed\nPage: %s\nWords: %v\n", page.Name, strings.Join(words.Words, ", "))
	case "occurrence":
		var occurrence Occurrence
		if err := json.Unmarshal(body, &occurrence); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("JSON Parsed\nPage: %s\nWords: %v\n", page.Name, occurrence.Words)
	default:
		fmt.Println("Page not found")
	}
}
