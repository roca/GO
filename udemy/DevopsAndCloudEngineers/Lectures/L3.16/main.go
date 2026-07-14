package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type IResponse interface {
	GetResponse() string
}

// {"page":"words","input":"word1","words":["word1"]}
type Page struct {
	Name string `json:"page"`
}

type Words struct {
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func (w Words) GetResponse() string {
	return fmt.Sprintf("%s", strings.Join(w.Words, ", "))
}

type Occurrence struct {
	Words map[string]int `json:"words"`
}

func (o Occurrence) GetResponse() string {
	out := []string{}
	for word, occurrence := range o.Words {
		out = append(out, fmt.Sprintf("%s (%d)", word, occurrence))
	}
	return fmt.Sprintf("%s", strings.Join(out, ", "))
}

func main() {
	var (
		requestURL string
		password   string
		parsedURL  *url.URL
		err        error
	)

	flag.StringVar(&requestURL, "url", "", "URL to request")
	flag.StringVar(&password, "password", "", "use a password to access our api")

	flag.Parse()

	if parsedURL, err = url.ParseRequestURI(requestURL); err != nil {
		fmt.Printf("URL Validation error: %s\n", err)
		flag.Usage()
		os.Exit(1)
	}

	client := http.Client{}

	if password != "" {
		token, err := doLoginRequest(client, parsedURL.Scheme+"://"+parsedURL.Host+"/login", password)
		if err != nil {
			if requestErr, ok := err.(RequestError); ok {
				fmt.Printf("Error: %s (HTTP Code: %d, Body:%s)\n", requestErr.Err, requestErr.HTTPCode, requestErr.Body)
				os.Exit(1)
			}
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		client.Transport = &MyJWTTransport{
			transport: http.DefaultTransport,
			token:     token,
		}
	}

	res, err := doRequest(client, parsedURL.String())
	if err != nil {
		if requestErr, ok := err.(RequestError); ok {
			fmt.Printf("Error: %s (HTTP Code: %d, Body:%s)\n", requestErr.Err, requestErr.HTTPCode, requestErr.Body)
			os.Exit(1)
		}
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	if res == nil {
		fmt.Println("No response")
		os.Exit(1)
	}

	fmt.Printf("Response: %s\n", res.GetResponse())
}

func doRequest(client http.Client, requestURL string) (IResponse, error) {

	response, err := client.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("http Get error: %s\n", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("ReadAll error: %s\n", err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Invalid output (HTTP Code %d): %s\n", response.StatusCode, body)
	}

	if !json.Valid(body) {
		return nil, RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(body),
			Err:      fmt.Sprintf("No valid JSON rerturned"),
		}
	}

	var page Page
	if err := json.Unmarshal(body, &page); err != nil {
		return nil, RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(body),
			Err:      fmt.Sprintf("Page unmarshal error: %s\n", err),
		}
	}

	switch page.Name {
	case "words":
		var words Words
		if err := json.Unmarshal(body, &words); err != nil {
			return nil, RequestError{
				HTTPCode: response.StatusCode,
				Body:     string(body),
				Err:      fmt.Sprintf("Words unmarshal error: %s\n", err),
			}
		}
		return words, nil
	case "occurrence":
		var occurrence Occurrence
		if err := json.Unmarshal(body, &occurrence); err != nil {
			return nil, RequestError{
				HTTPCode: response.StatusCode,
				Body:     string(body),
				Err:      fmt.Sprintf("Occurrence unmarshal error: %s\n", err),
			}
		}
		return occurrence, nil
	}

	return nil, fmt.Errorf("No responses for this url: %s", requestURL)
}
