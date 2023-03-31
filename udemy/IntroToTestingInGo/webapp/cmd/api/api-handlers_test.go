package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
	"webapp/pkg/data"
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

func Test_app_refresh(t *testing.T) {
	// test the refresh token endpoint

	var tests = []struct {
		name               string
		token              string
		expectedStatusCode int
		resetRefreshTime   bool
	}{
		{"valid token", "", http.StatusOK, true},
		{"valid token but not yet ready to expire", "", http.StatusTooEarly, false},
		{"expired token", expiredToken, http.StatusBadRequest, false},
	}

	testUser := data.User{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
		Email:     "admin@example.com",
	}

	oldRefreshTime := refreshTokenExpiry

	for _, e := range tests {
		t.Run(e.name, func(t *testing.T) {
			tkn := e.token
			if tkn == "" {
				if e.resetRefreshTime {
					refreshTokenExpiry = time.Second * 1
				}
				tokens, _ := app.generateTokenPair(&testUser)
				tkn = tokens.RefreshToken
			}

			postedData := url.Values{
				"refresh_token": {tkn},
			}
			req, _ := http.NewRequest("POST", "/refresh-token", strings.NewReader(postedData.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(app.refresh)
			handler.ServeHTTP(rr, req)

			if rr.Code != e.expectedStatusCode {
				t.Errorf("%s: expected status code %d, got %d", e.name, e.expectedStatusCode, rr.Code)
			}
		})
		refreshTokenExpiry = oldRefreshTime
	}

}
