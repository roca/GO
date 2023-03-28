package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"webapp/pkg/data"
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

func Test_app_authRequired(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	testUser := data.User{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
		Email:     "admin@example.com",
	}

	tokens, _ := app.generateTokenPair(&testUser)

	var tests = []struct {
		name            string
		token           string
		expectedAuthorized bool
		setHeader       bool
	}{
		{"valid token", fmt.Sprintf("Bearer %s", tokens.Token), true, true},
	}

	for _, e := range tests {
		t.Run(e.name, func(t *testing.T) {
			handlerToTest := app.authRequired(nextHandler)

			req, _ := http.NewRequest("GET", "http://", nil)
			if e.setHeader {
				req.Header.Set("Authorization", e.token)
			}
			rr := httptest.NewRecorder()
			handlerToTest.ServeHTTP(rr, req)

			if e.expectedAuthorized && rr.Code == http.StatusUnauthorized {
				t.Errorf("%s: got code 402, and should not have", e.name)
			}

			if !e.expectedAuthorized && rr.Code != http.StatusUnauthorized {
				t.Errorf("%s: Unexpected error: %d", e.name, rr.Code)
			}
		})
	}

}
