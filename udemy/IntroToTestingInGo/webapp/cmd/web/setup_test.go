package main

import (
	"os"
	"testing"
)

var app application

func TestMain(m *testing.M) {
	app.Session = getSession()
	
	code := m.Run()

	os.Exit(code)
}
