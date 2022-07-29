package api

import (
	"io"
	"net/http"
	"testing"
)

type MockClient struct {
	ResponseOutput *http.Response
}

func (m *MockClient) Get(url string) (*http.Response, error) {
	return m.ResponseOutput, nil
}

func TestDoGetRequest(t *testing.T) {
	apiInstance := api{
		Options: Options{},
		Client: &MockClient{
			ResponseOutput: &http.Response{
				StatusCode: 200,
				Body:       io.ReadCloser(nil),
			},
		},
	}

	resp, err := apiInstance.DoGetRequest("http://localhost:8080/words")
}
