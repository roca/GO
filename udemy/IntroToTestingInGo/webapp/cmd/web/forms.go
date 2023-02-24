package main

import (
	"net/url"
	"strings"
)

type errors map[string][]string

func (e errors) Get(field string) string {
	errorSlice := e[field]
	if len(errorSlice) == 0 {
		return ""
	}
	return errorSlice[0]
}

func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

type Form struct {
	Data   url.Values
	Errors errors
}

func NewForm(data url.Values) *Form {
	return &Form{
		Data:   data,
		Errors: errors(map[string][]string{}),
	}
}

func (f *Form) Has(field string) bool {
	x := f.Data.Get(field)
	return strings.TrimSpace(x) != ""
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		if !f.Has(field) {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

func (f *Form) Check(ok bool, field, message string) {
	if !ok {
		f.Errors.Add(field, message)
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
