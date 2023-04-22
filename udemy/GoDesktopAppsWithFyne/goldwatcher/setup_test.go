package main

import (
	"net/http"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

var jsonToReturn = `{
        "ts": 1682067394727,
        "tsj": 1682067385985,
        "date": "Apr 21st 2023, 04:56:25 am NY",
        "items": [
                {
                        "curr": "USD",
                        "xauPrice": 1988.0825,
                        "xagPrice": 25.1213,
                        "chgXau": -17.1825,
                        "chgXag": -0.1332,
                        "pcXau": -0.8569,
                        "pcXag": -0.5274,
                        "xauClose": 2005.265,
                        "xagClose": 25.2545
                }
        ]
}`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}
