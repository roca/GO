package toolkit_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"

	"github.com/roca/go-toolkit"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func TestTools_PushJSONToRemote(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body: ioutil.NopCloser(bytes.NewBufferString("ok")),
			Header: make(http.Header),
		}
	})

	var testTools toolkit.Tools
	var foo struct {
		Bar string `json:"bar"`
	}
	foo.Bar = "bar"

	_,_,err := testTools.PushJSONToRemote("http://localhost", foo,client)
	if err != nil {
		t.Errorf("PushJSONToRemote returned an error: %v", err)
	}
}

func TestTools_RandomString(t *testing.T) {
	var testTools toolkit.Tools
	s := testTools.RandomString(10)
	if len(s) != 10 {
		t.Errorf("RandomString returned a string of length %d, expected 10", len(s))
	}
}

var uploadTests = []struct {
	name          string
	allowedTypes  []string
	renameFile    bool
	errorExpected bool
}{
	{name: "allowed no rename", allowedTypes: []string{"image/jpeg", "image/png"}, renameFile: false, errorExpected: false},
	{name: "allowed rename", allowedTypes: []string{"image/jpeg", "image/png"}, renameFile: true, errorExpected: false},
	{name: "not rename", allowedTypes: []string{"image/jpeg"}, renameFile: false, errorExpected: true},
}

func TestTools_UploadFile(t *testing.T) {
	for _, e := range uploadTests {
		// set up a pipe to avoid buffering
		pr, pw := io.Pipe()
		writer := multipart.NewWriter(pw)
		wg := sync.WaitGroup{}
		wg.Add(1)

		go func() {
			defer writer.Close()
			defer wg.Done()

			// create the form data field 'file'
			part, err := writer.CreateFormFile("file", "./testdata/img.png")
			if err != nil {
				t.Errorf("CreateFormFile failed: %v", err)
			}

			f, err := os.Open("./testdata/img.png")
			if err != nil {
				t.Errorf("Open failed: %v", err)
			}
			defer f.Close()

			img, _, err := image.Decode(f)
			if err != nil {
				t.Errorf("Decode failed: %v", err)
			}

			err = png.Encode(part, img)
			if err != nil {
				t.Errorf("Encode failed: %v", err)
			}
		}()

		// read from the pipe which receives data
		request := httptest.NewRequest("POST", "/", pr)
		request.Header.Add("Content-Type", writer.FormDataContentType())

		var testTools toolkit.Tools
		testTools.AllowedFileTypes = e.allowedTypes

		uploadedFiles, err := testTools.UploadFiles(request, "./testdata/uploads", e.renameFile)
		if err != nil && !e.errorExpected {
			t.Errorf("UploadFile failed: %v", err)
		}

		if !e.errorExpected {
			if _, err := os.Stat(fmt.Sprintf("./testdata/uploads/%s", uploadedFiles[0].NewFileName)); os.IsNotExist(err) {
				t.Errorf("%s: expected file to exist: %s", e.name, err.Error())
			}

			// clean up
			_ = os.Remove(fmt.Sprintf("./testdata/uploads/%s", uploadedFiles[0].NewFileName))
		}

		if !e.errorExpected && err != nil {
			t.Errorf("%s: error expected but none reveived", e.name)
		}

		wg.Wait()

	}
}

func TestTools_UploadOneFile(t *testing.T) {
	// set up a pipe to avoid buffering
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)

	go func() {
		defer writer.Close()

		// create the form data field 'file'
		part, err := writer.CreateFormFile("file", "./testdata/img.png")
		if err != nil {
			t.Errorf("CreateFormFile failed: %v", err)
		}

		f, err := os.Open("./testdata/img.png")
		if err != nil {
			t.Errorf("Open failed: %v", err)
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			t.Errorf("Decode failed: %v", err)
		}

		err = png.Encode(part, img)
		if err != nil {
			t.Errorf("Encode failed: %v", err)
		}
	}()

	// read from the pipe which receives data
	request := httptest.NewRequest("POST", "/", pr)
	request.Header.Add("Content-Type", writer.FormDataContentType())

	var testTools toolkit.Tools

	uploadedFiles, err := testTools.UploadOneFile(request, "./testdata/uploads", true)
	if err != nil {
		t.Errorf("UploadFile failed: %v", err)
	}

	if _, err := os.Stat(fmt.Sprintf("./testdata/uploads/%s", uploadedFiles.NewFileName)); os.IsNotExist(err) {
		t.Errorf("expected file to exist: %s", err.Error())
	}

	// clean up
	_ = os.Remove(fmt.Sprintf("./testdata/uploads/%s", uploadedFiles.NewFileName))
}

func TestTools_CreateDirIfNotExist(t *testing.T) {
	var testTools toolkit.Tools
	path := "./testdata/mydir"
	err := testTools.CreateDirIfNotExist(path)
	if err != nil {
		t.Errorf("CreateDirIfNotExist failed: %v", err)
	}

	err = testTools.CreateDirIfNotExist(path)
	if err != nil {
		t.Errorf("CreateDirIfNotExist failed: %v", err)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("expected directory to exist: %s", err.Error())
	}

	os.RemoveAll(path)
}

var slugTests = []struct {
	name          string
	s             string
	expected      string
	errorExpected bool
}{
	{"valid string", "now is the time", "now-is-the-time", false},
	{"empty string", "", "", true},
	{"complex string", "Now is the time for all GOOD men! + fish & such &^123", "now-is-the-time-for-all-good-men-fish-such-123", false},
	{"japanese string", "こんにちは世界", "", true},
	{"japanese string and roman characters", "hello こんにちは世界 world", "hello-world", false},
}

func TestTools_Slugify(t *testing.T) {
	var testTools toolkit.Tools
	for _, e := range slugTests {
		t.Run(e.name, func(t *testing.T) {
			slug, err := testTools.Slugify(e.s)
			if err != nil && !e.errorExpected {
				t.Errorf("%s: error received when none expected: %s", e.name, err.Error())
			}
			if !e.errorExpected && slug != e.expected {
				t.Errorf("%s: wrong slug returned; expected %s but got %s", e.name, e.expected, slug)
			}
		})
	}

}

func TestTools_DownloadStaticFile(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	tools := toolkit.Tools{}
	tools.DownloadStaticFile(rr, req, "./testdata", "pic.jpg", "puppy.jpeg")

	res := rr.Result()
	defer res.Body.Close()

	if res.Header["Content-Length"][0] != "98827" {
		t.Errorf("expected Content-Length to be %s but got %s", "114", res.Header["Content-Length"][0])
	}

	if res.Header["Content-Disposition"][0] != "attachment; filename=\"puppy.jpeg\"" {
		t.Errorf("expected Content-Disposition to be %s but got %s", "attachment; filename=puppy.jpeg", res.Header["Content-Disposition"][0])
	}

	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("ReadAll failed: %v", err)
	}
}

var jsonTests = []struct {
	name          string
	json          string
	errorExpected bool
	maxSize       int
	allowUnknown  bool
}{
	{name: "good json", json: `{"foo": "bar"}`, errorExpected: false, maxSize: 1024, allowUnknown: false},
	{name: "badly formatted json", json: `{"foo": }`, errorExpected: true, maxSize: 1024, allowUnknown: false},
	{name: "incorrect type", json: `{"foo": 1}`, errorExpected: true, maxSize: 1024, allowUnknown: false},
	{name: "two json files", json: `{"foo": "1"}{"alpha": "beta"}`, errorExpected: true, maxSize: 1024, allowUnknown: false},
	{name: "empty body", json: ``, errorExpected: true, maxSize: 1024, allowUnknown: false},
	{name: "syntax error in json", json: `{"foo": 1"`, errorExpected: true, maxSize: 1024, allowUnknown: false},
	{name: "unknown field in json", json: `{"food": "1"}`, errorExpected: true, maxSize: 1024, allowUnknown: false},
	{name: "allow unknown fields in json", json: `{"food": "1"}`, errorExpected: false, maxSize: 1024, allowUnknown: true},
	{name: "missing field name", json: `{jack: "1"}`, errorExpected: true, maxSize: 1024, allowUnknown: true},
	{name: "file too large", json: `{"foo": "bar"}`, errorExpected: true, maxSize: 5, allowUnknown: true},
	{name: "not json", json: `hello world`, errorExpected: true, maxSize: 1024, allowUnknown: true},
}

func TestTools_ReadJSON(t *testing.T) {
	var testTool toolkit.Tools
	for _, e := range jsonTests {
		t.Run(e.name, func(t *testing.T) {
			// set the max file size
			testTool.MaxJSONSize = e.maxSize

			// allow/disallow unknown fields
			testTool.AllowUnknownFields = e.allowUnknown

			// declare a variable to read the decoded json into
			var decodedJSON struct {
				Foo string `json:"foo"`
			}

			// creat a request with the json
			req := httptest.NewRequest("POST", "/", bytes.NewReader([]byte(e.json)))
			// create a response recorder
			rr := httptest.NewRecorder()

			// read the json
			err := testTool.ReadJSON(rr, req, &decodedJSON)

			if e.errorExpected && err == nil {
				t.Errorf("%s: no error received when one expected", e.name)
			}

			if !e.errorExpected && err != nil {
				t.Errorf("%s: error received when none expected: %s", e.name, err.Error())
			}

			req.Body.Close()
		})
	}
}

func TestTools_WriteJSON(t *testing.T) {
	var testTool toolkit.Tools

	rr := httptest.NewRecorder()
	payload := toolkit.JSONResponse{
		Error:   false,
		Message: "foo",
	}

	headers := make(http.Header)
	headers.Add("FOO", "BAR")

	err := testTool.WriteJSON(rr, http.StatusOK, payload, headers)
	if err != nil {
		t.Errorf("failed to write JSON %v", err)
	}
}

func TestTools_ErrorJSON(t *testing.T) {
	var testTool toolkit.Tools

	rr := httptest.NewRecorder()
	err := testTool.ErrorJSON(rr, errors.New("some error"), http.StatusServiceUnavailable)
	if err != nil {
		t.Errorf("failed to write JSON %v", err)
	}

	var payload toolkit.JSONResponse
	decoder := json.NewDecoder(rr.Body)
	err = decoder.Decode(&payload)
	if err != nil {
		t.Errorf("failed to decode JSON %v", err)
	}

	if !payload.Error {
		t.Errorf("expected error to be true")
	}

	if rr.Code != http.StatusServiceUnavailable {
		t.Errorf("expected status code to be %d but got %d", http.StatusServiceUnavailable, rr.Code)
	}

	if payload.Message != "some error" {
		t.Errorf("expected message to be %s but got %s", "some error", payload.Message)
	}

}
