//go:build !integration

package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestListAction(t *testing.T) {
	testCases := []struct {
		name     string
		expError error
		expOut   string
		resp     struct {
			Status int
			Body   string
		}
		closeServer bool
	}{
		{name: "Results",
			expError: nil,
			expOut:   "-  1  Task 1\n-  2  Task 2\n",
			resp:     testResp["resultMany"],
		},
		{name: "NoResults",
			expError: ErrNotFound,
			resp:     testResp["noResults"],
		},
		{name: "InvalidURL",
			expError:    ErrConnection,
			resp:        testResp["noResults"],
			closeServer: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url, cleanup := mockServer(
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(tc.resp.Status)
					fmt.Fprint(w, tc.resp.Body)
				})
			defer cleanup()

			if tc.closeServer {
				cleanup()
			}

			var out bytes.Buffer

			err := listAction(&out, url)

			if tc.expError != nil {
				if err == nil {
					t.Errorf("Expected error: %q, go no error", tc.expError)
				}

				if !errors.Is(err, tc.expError) {
					t.Errorf("Expected error: %q, got: %q", tc.expError, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("Expected no error, got %q", err)
			}

			if tc.expOut != out.String() {
				t.Errorf("Expected output: %q, got: %q", tc.expOut, out.String())
			}

		})
	}
}

func TestViewAction(t *testing.T) {
	testCases := []struct {
		name     string
		expError error
		expOut   string
		resp     struct {
			Status int
			Body   string
		}
		id string
	}{
		{name: "ResultsOne",
			expError: nil,
			expOut:   "Task:         Task 1\nCreated at:   Oct/28 @08:23\nCompleted:    No\n",
			resp:     testResp["resultOne"],
			id:       "1",
		},
		{name: "NotFound",
			expError: ErrNotFound,
			resp:     testResp["notFound"],
			id:       "1",
		},
		{name: "InvalidID",
			expError: ErrNotNumber,
			resp:     testResp["notResults"],
			id:       "a",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url, cleanup := mockServer(
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(tc.resp.Status)
					fmt.Fprintln(w, tc.resp.Body)
				})
			defer cleanup()

			var out bytes.Buffer

			err := viewAction(&out, url, tc.id)

			if tc.expError != nil {
				if err == nil {
					t.Errorf("Expected error: %q, go no error", tc.expError)
				}

				if !errors.Is(err, tc.expError) {
					t.Errorf("Expected error: %q, got: %q", tc.expError, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("Expected no error, got %q", err)
			}

			if tc.expOut != out.String() {
				t.Errorf("Expected output: %q, got: %q", tc.expOut, out.String())
			}
		})
	}
}

func TestAddAction(t *testing.T) {
	expURLPath := "/todo"
	expMethod := http.MethodPost
	expBody := "{\"task\":\"Task 1\"}\n"
	expContentType := "application/json"
	expOut := "Added task \"Task 1\" to your list.\n"
	args := []string{"Task", "1"}

	url, cleanup := mockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != expURLPath {
			t.Errorf("Expected URL path: %q, got: %q", expURLPath, r.URL.Path)
		}

		if r.Method != expMethod {
			t.Errorf("Expected method: %q, got: %q", expMethod, r.Method)
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		r.Body.Close()

		if string(body) != expBody {
			t.Errorf("Expected body: %q, got: %q", expBody, string(body))
		}

		if r.Header.Get("Content-Type") != expContentType {
			t.Errorf("Expected Content-Type: %q, got: %q", expContentType, r.Header.Get("Content-Type"))
		}
		w.WriteHeader(testResp["created"].Status)
		fmt.Fprintln(w, testResp["created"].Body)
	})
	defer cleanup()

	var out bytes.Buffer

	if err := addAction(&out, url, args); err != nil {
		t.Fatalf("Expected no error, got %q", err)
	}

	if expOut != out.String() {
		t.Errorf("Expected output: %q, got: %q", expOut, out.String())
	}
}

func TestCompleteAction(t *testing.T) {
	expURLPath := "/todo/1"
	expMethod := http.MethodPatch
	expQuery := "complete"
	expOut := "Item number 1 marked as completed.\n"
	arg := "1"
	url, cleanup := mockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != expURLPath {
			t.Errorf("Expected URL path: %q, got: %q", expURLPath, r.URL.Path)
		}

		if r.Method != expMethod {
			t.Errorf("Expected method: %q, got: %q", expMethod, r.Method)
		}

		if _, ok := r.URL.Query()[expQuery]; !ok {
			t.Errorf("Expected query: %q, got: %q", expQuery, r.URL.Query())
		}

		w.WriteHeader(testResp["noContent"].Status)
		fmt.Fprintln(w, testResp["noContent0"].Body)
	})
	defer cleanup()

	var out bytes.Buffer

	if err := completeAction(&out, url, arg); err != nil {
		t.Fatalf("Expected no error, got %q", err)
	}

	if expOut != out.String() {
		t.Errorf("Expected output: %q, got: %q", expOut, out.String())
	}
}

func TestDelAction(t *testing.T) {
	expURLPath := "/todo/1"
	expMethod := http.MethodDelete
	expOut := "Item number 1 deleted.\n"
	arg := "1"

	url, cleanup := mockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != expURLPath {
			t.Errorf("Expected URL path: %q, got: %q", expURLPath, r.URL.Path)
		}

		if r.Method != expMethod {
			t.Errorf("Expected method: %q, got: %q", expMethod, r.Method)
		}

		w.WriteHeader(testResp["noContent"].Status)
		fmt.Fprintln(w, testResp["noContent"].Body)
	})
	defer cleanup()

	var out bytes.Buffer

	if err := delAction(&out, url, arg); err != nil {
		t.Fatalf("Expected no error, got %q", err)
	}

	if expOut != out.String() {
		t.Errorf("Expected output: %q, got: %q", expOut, out.String())
	}
}
