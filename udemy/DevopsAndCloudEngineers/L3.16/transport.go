package main

import "net/http"

type MyJWTTransport struct {
	transport http.RoundTripper
	token     string
}

func (m *MyJWTTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+m.token)
	return m.transport.RoundTrip(req)
}
