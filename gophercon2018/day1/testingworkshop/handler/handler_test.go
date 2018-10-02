package handler_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/GOCODE/gophercon2018/day1/testingworkshop/handler"
)

func Test_U_FormHandler_Template_Error(t *testing.T) {

	form := url.Values{}
	form.Add("name", "John")

	// pass invalid hex strings
	// The `Encode` method on `url.Values` will properly encode the values we set into well formed `string` that can be read as the body of the request.
	req := httptest.NewRequest("POST", "/form", strings.NewReader(form.Encode()))

	// We must set the `Content-Type` correctly for `ParseForm` to work.
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// set the header `Content-Type` to `application/x-www-form-urlencoded`
	// create a new httptest.NewRecorder
	res := httptest.NewRecorder()

	// Call the FormHandler
	// Test to see the the response code is 500

	handler.FormHandler(res, req)

	if got, exp := res.Code, http.StatusOK; got != exp {
		t.Errorf("unexpected response code.  got: %d, exp %d\n", got, exp)
	}
	if got, exp := res.Body.String(), "Posted Hello, John!"; got != exp {
		t.Errorf("unexpected body.  got: %s, exp %s\n", got, exp)
	}

	// test the body is `Oops!`
}
