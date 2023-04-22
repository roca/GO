package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_client(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
			Header:     make(http.Header),
		}
	})

	req, _ := http.NewRequest("GET", "http://localhost", nil)

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Test client error %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Test client body\n%v: error %s", string(body), err)
	}

	gold := Gold{}
	err = json.Unmarshal(body, &gold)
	if err != nil {
		t.Errorf("Test client unmarshal error %s\nBody: %s\nGold object: %v", err, string(body), gold)
	}
}

func TestGold_GetPrices(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
			Header:     make(http.Header),
		}
	})

	g := Gold{
		Prices: nil,
		Client: client,
	}

	p, err := g.GetPrices()
	if err != nil {
		t.Error(err)
	}

	if p.Price != 1988.0825 {
		t.Errorf("Wrong price returned: expected %f, got %f", 1988.0825, p.Price)
	}
}
