package main

import (
	"os"
	"testing"
)

var app application

func TestMain(m *testing.M) {
	app = application{}
	
	code := m.Run()

	os.Exit(code)
}
