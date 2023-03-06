package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"webapp/pkg/data"
)

func Test_app_addIPToContest(t *testing.T) {
	tests := []struct {
		headerName  string
		headerValue string
		addr        string
		emptyAddr   bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "192.3.2.1", "", false},
		{"", "", "hello:world", false},
	}

	// create a dummy handler that will check the context
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Error(contextUserKey, "is not present")
		}
		ip, ok := val.(string)
		if !ok {
			t.Error("context value is not a string")
		}
		t.Log("ip:", ip)
	})

	for _, e := range tests {
		handlerToTest := app.addIPToContext(nextHandler)

		req := httptest.NewRequest("GET", "http://testing", nil)

		if e.emptyAddr {
			req.RemoteAddr = ""
		}

		if len(e.headerName) > 0 {
			req.Header.Add(e.headerName, e.headerValue)
		}

		if len(e.addr) > 0 {
			req.RemoteAddr = e.addr
		}

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}

func Test_app_ipFromContext(t *testing.T) {

	ctx := context.WithValue(context.Background(), contextUserKey, "some_ip")
	ip := app.ipFromContext(ctx)
	if !strings.EqualFold(ip, "some_ip") {
		t.Error("ipFromContext returned wrong ip")
	}
}

func Test_app_auth(t *testing.T) {
	tests := []struct {
		namespace string
		isAuth    bool
	}{
		{"logged in ", true},
		{"not logged in", false},
	}

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

	for _, e := range tests {
		handlerToTest := app.auth(nextHandler)
		req := httptest.NewRequest("GET", "http://testing", nil)
		req = addContextAndSessionToRequest(req, app)
		if e.isAuth {
			app.Session.Put(req.Context(), "user", data.User{ID: 1})
		}
		rr := httptest.NewRecorder()
		handlerToTest.ServeHTTP(rr, req)

		if e.isAuth && rr.Code != http.StatusOK {
			t.Errorf("%s: expected status code %d, got %d", e.namespace, http.StatusOK, rr.Code)
		}

		if !e.isAuth && rr.Code != http.StatusTemporaryRedirect {
			t.Errorf("%s: expected status code %d, got %d", e.namespace, http.StatusTemporaryRedirect, rr.Code)
		}

	}
}
