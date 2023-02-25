package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Has(t *testing.T) {
	form := NewForm(nil)

	has := form.Has("foo")
	if has {
		t.Error("form show has field when is should not")
	}

	postedData := url.Values{}
	postedData.Add("foo", "bar")
	form = NewForm(postedData)

	has = form.Has("foo")
	if !has {
		t.Error("shows form dose not have field when it should")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/foo", nil)
	form := NewForm(r.PostForm)

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("form shows valid when it should not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/foo", nil)
	r.PostForm = postedData

	form = NewForm(r.PostForm)
	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("form shows invalid when it should not")
	}
}

func TestForm_Check(t *testing.T) {
	form := NewForm(nil)

	form.Check(false,"password", "password is required")
	if form.Valid() {
		t.Error("Valid() returns false, and it should be true when calling Check()")
	}
}

func TestForm_ErrorGet(t *testing.T) {
	form := NewForm(nil)
	form.Check(false,"password", "password is required")
	s := form.Errors.Get("password")

	if s != "password is required" {
		t.Error("Error.Get() returns wrong error message")
	}

	s = form.Errors.Get("foo")
	if s != "" {
		t.Error("Error.Get() returns error message when it should not")
	}
}
