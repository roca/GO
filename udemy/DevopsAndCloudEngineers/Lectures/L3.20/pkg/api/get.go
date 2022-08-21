package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

type WordPage struct {
	Page
	Words
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

type Assignment struct {
	Page         string             `json:"page"`
	Words        []string           `json:"words"`
	Percentages  map[string]float64 `json:"percentages"`
	Special      []interface{}      `json:"special"`
	ExtraSpecial []interface{}      `json:"extra_special"`
}

func (a Assignment) GetResponse() string {
	asBytes, err := json.MarshalIndent(a, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	return string(asBytes)
}


func (a api) DoGetRequest(requestURL string) (IResponse, error) {

	response, err := a.Client.Get(requestURL)
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
	case "assignment":
		var assignment Assignment
		if err := json.Unmarshal(body, &assignment); err != nil {
			return nil, RequestError{
				HTTPCode: response.StatusCode,
				Body:     string(body),
				Err:      fmt.Sprintf("Assignment unmarshal error: %s\n", err),
			}
		}
		return assignment, nil
	}

	return nil, fmt.Errorf("No responses for this url: %s", requestURL)
}
