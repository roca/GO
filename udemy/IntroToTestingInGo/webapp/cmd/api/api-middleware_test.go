package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_app_enableCORS(t *testing.T) {

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	var test = []struct {
		name         string
		method       string
		expectHeader bool
	}{
		{"pre-flight", "OPTIONS", true},
		{"get", "GET", false},
	}

	for _, e := range test {
		t.Run(e.name, func(t *testing.T) {
			handlerToTest := app.enableCORS(nextHandler)

			req := httptest.NewRequest(e.method, "http://testing", nil)
			rr := httptest.NewRecorder()

			handlerToTest.ServeHTTP(rr, req)

			if e.expectHeader && rr.Header().Get("Access-Control-Allow-Credentials") == "" {
				t.Errorf("%s: expected header %s to be set", e.name, "Access-Control-Allow-Credentials")
			}

			if !e.expectHeader && rr.Header().Get("Access-Control-Allow-Credentials") != "" {
				t.Errorf("%s: expected header %s to be empty", e.name, "Access-Control-Allow-Credentials")
			}

		})
	}
}
