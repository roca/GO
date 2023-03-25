package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_app_authenticate(t *testing.T) {
	var theTests = []struct {
		name               string
		requestBody        string
		expectedStatusCode int
	}{
		{"valid user", `{"email": "admin@example.com","password":"secret"}`, http.StatusOK},
		{"not json", `I'm not JSON`, http.StatusUnauthorized},
		{"empty json", `{}`, http.StatusUnauthorized},
		{"empty email", `{"email": "","password":"secret"}`, http.StatusUnauthorized},
		{"no password", `{"email": "admin@example.com","password":""}`, http.StatusUnauthorized},
		{"invalid user", `{"email": "admin@someotherdomain.com","password":"secret"}`, http.StatusUnauthorized},
	}

	for _, e := range theTests {
		t.Run(e.name, func(t *testing.T) {
			var reader io.Reader
			reader = strings.NewReader(e.requestBody)
			req, _ := http.NewRequest("POST", "/auth", reader)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(app.authenticate)

			handler.ServeHTTP(rr, req)

			if e.expectedStatusCode != rr.Code {
				t.Errorf("%s: Expected status code %d, got %d", e.name, e.expectedStatusCode, rr.Code)
			}
		})
	}
}
