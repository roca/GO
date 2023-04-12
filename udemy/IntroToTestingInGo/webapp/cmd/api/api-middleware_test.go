package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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
		name               string
		token              string
		expectedAuthorized bool
		setHeader          bool
	}{
		{"valid token", fmt.Sprintf("Bearer %s", tokens.Token), true, true},
		{"no token", "", false, false},
		{"invalid token", fmt.Sprintf("Bearer %s", expiredToken), false, true},
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
				t.Errorf("%s: got code 401, and should not have", e.name)
			}

			if !e.expectedAuthorized && rr.Code != http.StatusUnauthorized {
				t.Errorf("%s: Unexpected error: %d", e.name, rr.Code)
			}
		})
	}

}

func Test_app_refreshUsingCookie(t *testing.T) {
	testUser := data.User{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
		Email:     "admin@example.com",
	}
	tokens, _ := app.generateTokenPair(&testUser)

	testCookie := &http.Cookie{
		Name:     "_Host-refresh_token",
		Path:     "/",
		Value:    tokens.RefreshToken,
		Expires:  time.Now().Add(refreshTokenExpiry),
		MaxAge:   int(refreshTokenExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		Domain:   "localhost",
		HttpOnly: true,
		Secure:   true,
	}

	badCookie := &http.Cookie{
		Name:     "_Host-refresh_token",
		Path:     "/",
		Value:    "bad-token",
		Expires:  time.Now().Add(refreshTokenExpiry),
		MaxAge:   int(refreshTokenExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		Domain:   "localhost",
		HttpOnly: true,
		Secure:   true,
	}

	var tests = []struct {
		name               string
		addCookie          bool
		cookie             *http.Cookie
		expectedStatusCode int
	}{
		{"valid cookie", true, testCookie, http.StatusOK},
		{"invalid cookie", true, badCookie, http.StatusBadRequest},
		{"no cookie", false, nil, http.StatusUnauthorized},
	}

	for _, e := range tests {
		t.Run(e.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			req, _ := http.NewRequest("GET", "http://", nil)
			if e.addCookie {
				req.AddCookie(e.cookie)
			}

			handler := http.HandlerFunc(app.refreshUsingCookie)
			handler.ServeHTTP(rr, req)

			if rr.Code != e.expectedStatusCode {
				t.Errorf("%s: Expected status code %d, got %d", e.name, e.expectedStatusCode, rr.Code)
			}

		})
	}
}

func Test_app_deleteRefreshCookie(t *testing.T) {
	req, _ := http.NewRequest("GET", "/logout", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.deleteRefreshCookie)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusAccepted {
		t.Errorf("Expected status code %d, got %d", http.StatusAccepted, rr.Code)
	}

	foundCookie := false
	for _, c := range rr.Result().Cookies() {
		if c.Name == "_Host-refresh_token" {
			foundCookie = true
			if c.Expires.After(time.Now()) {
				t.Errorf("Expected cookie to be expired, but it was not %v", c.Expires.UTC())
			}
		}
	}

	if !foundCookie {
		t.Errorf("Expected to find cookie, but did not")
	}
}
