package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	theTests := []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},
		{"404", "/notfound", http.StatusNotFound},
	}

	routes := app.routes()

	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	pathToTemplates = "./../../templates/"

	for _, e := range theTests {
		t.Run(e.name, func(t *testing.T) {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("want %d; got %d", e.expectedStatusCode, resp.StatusCode)
			}
		})
	}
}

// func TestAppHomeOld(t *testing.T) {
// 	req, _ := http.NewRequest("GET", "/", nil)
// 	req = addContextAndSessionToRequest(req, app)

// 	rr := httptest.NewRecorder()

// 	handler := http.HandlerFunc(app.Home)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusOK {
// 		t.Errorf("TestAppHome returned wrong status code; expected %d but got %d", http.StatusOK, rr.Code)
// 	}

// 	body, _ := io.ReadAll(rr.Body)

// 	if  !strings.Contains(string(body), `<small>From Session:`) {
// 		t.Error("TestAppHome returned unexpected body")
// 	}
// }

func TestAppHome(t *testing.T) {
	var tests = []struct{
		name string
		putInSession string
		expectedHTML string
	}{
		{"first visit", "", `<small>From Session:`},
		{"second visit", "hello, world", `<small>From Session: hello, world`},
	}

	for _, e := range tests {
		t.Run(e.name, func(t *testing.T) {
			pathToTemplates = "./../../templates/"
			req, _ := http.NewRequest("GET", "/", nil)
			req = addContextAndSessionToRequest(req, app)
			_ = app.Session.Destroy(req.Context())

			if e.putInSession != "" {
				app.Session.Put(req.Context(), "test", e.putInSession)
			}

			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(app.Home)

			handler.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK {
				t.Errorf("TestAppHome returned wrong status code; expected %d but got %d", http.StatusOK, rr.Code)
			}

			body, _ := io.ReadAll(rr.Body)

			if  !strings.Contains(string(body), e.expectedHTML) {
				t.Errorf("%s: did not find %s in response body", e.name, e.expectedHTML)
			}
		})
	}
}

func getCtx(req *http.Request) context.Context {
	ctx := context.WithValue(req.Context(), contextUserKey, "unknown")
	return ctx
}

func addContextAndSessionToRequest(req *http.Request, app application) *http.Request {
	req = req.WithContext(getCtx(req))
	ctx, _ := app.Session.Load(req.Context(), req.Header.Get("X-Session"))

	return req.WithContext(ctx)
}
