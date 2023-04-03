package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
	"webapp/pkg/data"

	"github.com/go-chi/chi/v5"
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

func Test_app_userHandlers(t *testing.T) {
	var tests = []struct {
		name               string
		method             string
		json               string
		paramID            string
		handler            http.HandlerFunc
		expectedStatusCode int
	}{
		{"allUsers", "GET", "", "", app.allUser, http.StatusOK},
		{"deleteUsers", "DELETE", "", "1", app.deleteUser, http.StatusNoContent},
		{"deleteUsers bad URL param", "DELETE", "", "Y", app.deleteUser, http.StatusBadRequest},
		{"getUser valid", "GET", "", "1", app.getUser, http.StatusOK},
		{"getUser invalid", "GET", "", "100", app.getUser, http.StatusBadRequest},
		{"getUser bad URL param", "GET", "", "Y", app.getUser, http.StatusBadRequest},
		{
			"updateUser valid",
			"PATCH",
			`{"id":1,"first_name":"Administrator","last_name":"User","email":"admin@example.com"}`,
			"",
			app.updateUser,
			http.StatusNoContent,
		},
		{
			"updateUser invalid",
			"PATCH",
			`{"id":100,"first_name":"Administrator","last_name":"User","email":"admin@example.com"}`,
			"",
			app.updateUser,
			http.StatusBadRequest,
		},
		{
			"insertUser valid",
			"PUT",
			`{"first_name":"Jack","last_name":"Smith","email":"Jack@example.com"}`,
			"",
			app.insertUser,
			http.StatusNoContent,
		},
		{
			"insertUser invalid",
			"PUT",
			`{"foo": "bar", "first_name":"Jack","last_name":"Smith","email":"Jack@example.com"}`,
			"",
			app.insertUser,
			http.StatusBadRequest,
		},
		{
			"insertUser invalid json",
			"PUT",
			`{first_name":"Jack","last_name":"Smith","email":"Jack@example.com"}`,
			"",
			app.insertUser,
			http.StatusBadRequest,
		},
	}

	for _, e := range tests {
		t.Run(e.name, func(t *testing.T) {
			var req *http.Request
			if e.json == "" {
				req, _ = http.NewRequest(e.method, "/", nil)
			} else {
				req, _ = http.NewRequest(e.method, "/", strings.NewReader(e.json))
			}

			if e.paramID != "" {
				chiCtx := chi.NewRouteContext()
				chiCtx.URLParams.Add("userID", e.paramID)
				req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(e.handler)
			handler.ServeHTTP(rr, req)

			if rr.Code != e.expectedStatusCode {
				t.Errorf("%s: expected status code %d, got %d", e.name, e.expectedStatusCode, rr.Code)
			}
		})
	}
}
