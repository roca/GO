package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

type MockClient struct {
	GetResponseOutput     *http.Response
	PostResponseOutput *http.Response
}

func (m *MockClient) Get(url string) (*http.Response, error) {
	return m.GetResponseOutput, nil
}

func (m *MockClient) Post(url string, contenType string, body io.Reader) (resp *http.Response, err error) {
	return m.PostResponseOutput, nil
}

func TestDoGetRequest(t *testing.T) {
	wordPage := WordPage{
		Page: Page{"words"},
		Words: Words{
			Input: "abc",
			Words: []string{"a", "b"},
		},
	}
	wordPageBytes, err := json.Marshal(wordPage)
	if err != nil {
		t.Errorf("Error marshalling wordPage: %s", err)
	}

	apiInstance := api{
		Options: Options{},
		Client: &MockClient{
			GetResponseOutput: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewBuffer(wordPageBytes)),
			},
		},
	}

	response, err := apiInstance.DoGetRequest("http://localhost:8080/words")
	if err != nil {
		t.Errorf("Error DoingGetRequest: %s", err)
	}
	if response == nil {
		t.Fatalf("Response is empty")
	}
	if response.GetResponse() != "a, b" {
		t.Errorf("Invalid response: %s", response.GetResponse())
	}
}
